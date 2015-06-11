from waferslim.converters import converter_for,convert_arg

import boto
from boto.s3.key import Key

class S3Buckets(object):
    
    def query(self):

        s3 = boto.connect_s3()

        bs = s3.get_all_buckets()
        res = [[['Name',x.name]] for x in bs]

        converter = converter_for(list)
        return converter.to_string(res)

class S3Logs(object):

    def __init__(self, bucket, prefix):

        self.prefix = prefix
	self.bucket = bucket

    def query(self):
        
	s3 = boto.connect_s3()
	b = s3.get_bucket(self.bucket)

	res = [[['Path',x.key]] for x in b.list(self.prefix)]

        converter = converter_for(list)
	return converter.to_string(res)

class Logs(object):

    @convert_arg(to_type=int)
    def __init__(self, num):

        self.keys = [
            "CustomerA/MachineB/Agent%s/MAgE_20150331_023936.zip" % i 
            for i in range(0,num)
        ]

        print self.keys

    def clean_s3_up(self):

        s3 = boto.connect_s3()
        b = s3.get_bucket('dmp-log-analysis')

        [b.delete_key(k) for k in self.keys]   
    
    def upload_logs_to_s3(self):

        s3 = boto.connect_s3()
        b = s3.get_bucket('dmp-log-analysis')

        for k in self.keys:

            f1 = Key(b,k)
            f1.set_contents_from_filename(
                '/data/MAgE_20150331_023936.log'
            );

    def logs_on_s3(self):

       s3 = boto.connect_s3()
       b = s3.get_bucket('dmp-log-analysis')

       return len([k for k in b.list('CustomerA')])
    

