package main

import (
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type getFolderFunc func(name,
	folder,
	delim,
	marker string) *s3.ListResp

type bucket struct {
	name  string
	wFdrs *sync.WaitGroup

	getFolderFunc
}

func (b *bucket) get(folder,
	delim,
	marker string) *s3.ListResp {
	return b.getFolderFunc(b.name,
		folder,
		delim,
		marker)
}

type folder struct {
	bucket
	folder *s3.ListResp
}

type T interface{}

var auth = make(chan aws.Auth)

func authGen() {
	for {
		res, _ := aws.GetAuth("", "")
		auth <- res
	}
}

func retryCall(
	f func() (T, error)) T {

	type Res struct {
		t   T
		err error
	}

	for {

		res := make(chan Res)

		go func() {

			objs, err := f()
			res <- Res{objs, err}
		}()

		select {

		case r := <-res:

			if r.err == nil {
				return r.t
			}

			glog.Info(r.err)

		case <-time.After(time.Second * 30):

			glog.Info("Timeout, retrying...")

		}

	}

}

func getFolder(name,
	folder,
	delim,
	marker string) *s3.ListResp {

	s := s3.New(<-auth, aws.USEast)
	b := s.Bucket(name)

	f := func() (T, error) {
		return b.List(folder, delim, marker, 1000)
	}

	return retryCall(f).(*s3.ListResp)

}

func folderItems(f folder, items chan s3.Key) {

	lastKey := ""

	for _, o := range f.folder.Contents {

		items <- o
		lastKey = o.Key

	}

	if f.folder.IsTruncated {

		marker := f.folder.NextMarker

		if marker == "" {
			marker = lastKey
		}

		objs := f.bucket.get(f.folder.Prefix,
			"",
			marker)

		go folderItems(folder{f.bucket, objs}, items)

	} else {
		f.bucket.wFdrs.Done()
	}

}

func enumItems() (fldrs chan folder, keys chan s3.Key) {

	fldrs = make(chan folder)
	keys = make(chan s3.Key)

	go func() {
		for f := range fldrs {
			go folderItems(f, keys)
		}
	}()

	return

}

func listFolders(b bucket,
	parent string,
	folders chan<- folder) {

	objs := b.get(parent, "/", "")

	num := len(objs.CommonPrefixes)
	b.wFdrs.Add(num)

	for _, v := range objs.CommonPrefixes {

		go func(p string) {
			listFolders(b, p, folders)
		}(v)

	}

	folders <- folder{b, objs}

}

func walkBucket(name,
	parent string,
	get getFolderFunc) chan s3.Key {

	go authGen()

	var wFdrs sync.WaitGroup
	bucket := bucket{name, &wFdrs, get}

	wFdrs.Add(1)

	folders, items := enumItems()
	go listFolders(bucket, parent, folders)

	go func() {

		wFdrs.Wait()
		close(items)

	}()

	return items

}
