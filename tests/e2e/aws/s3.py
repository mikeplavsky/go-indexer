from waferslim.converters import converter_for,convert_arg

import boto
from boto.s3.key import Key

class Ssh(object):

    def __init__(self):

	import spur

        self.ssh = spur.SshShell(
	    hostname='localhost',
	    username='ec2-user',
	    private_key_file='/data/mp.pem',
	    missing_host_key=spur.ssh.MissingHostKey.accept
	)

    def get(self, url):
        
	try:

	   res = self.ssh.run([
	       'curl',
	       '-XGET',
	       url
	   ])

	except Exception as e:
	   return e

	return res.output   

    def delete(self, url):
        
	try:

	   self.ssh.run([
	       'curl',
	       '-XDELETE',
	       url
	   ])

	except Exception as e:
	   return e

	return True   


    def remove_container(self, name):
        
	try:

	    self.ssh.run([
	        'docker',
	        'rm',
	        '-f',
	         name])

	except Exception as e:
	    return e

	return True

    def run_script(self,name):

        try:

	   res = self.ssh.run([
	       name
	   ],
	   cwd='/home/ec2-user/go-indexer/')

	except Exception as e:
	   return e

        print(res.output)
        return True	   

    def run_script_for(self,name,customer):

        try:

	   res = self.ssh.run([
	       name
	   ],
           update_env=dict(
               AWS_CUSTOMER=customer
           ),
	   cwd='/home/ec2-user/go-indexer/')

	except Exception as e:
	   return e

        print(res.output)
        return True	   

class DeadLetterQueue(object):

    def __init__(self,instance_id):
        
	self.instance_id = instance_id

	import boto
	sqs = boto.connect_sqs()

	self.queue = sqs.get_queue(
	    self.instance_id + "_dl") 

    def get_count(self):
	return self.queue.count() 

    @convert_arg(to_type=int)
    def wait_seconds_till_count_increments_on(self, seconds, prev_count, num):

	count = self.queue.count()

	print count
	print prev_count
	print count < prev_count

	if count < prev_count + num:

	   import time
	   time.sleep(seconds)

	   return prev_count
	
	return count 


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

    def __init__(self, customer):
        self.customer = customer

    def clean_s3_up(self):

        s3 = boto.connect_s3()
        b = s3.get_bucket('dmp-log-analysis')
	
	keys = [x.key for x in b.list(self.customer)]
        
        [b.delete_key(k) for k in keys]   
    
    @convert_arg(to_type=int)
    def upload_logs_to_s3(self, num):

        s3 = boto.connect_s3()
        b = s3.get_bucket('dmp-log-analysis')

        for i in range(0,num):
            
            k = self.customer + "/MachineB/Agent%s/MAgE_20150331_023936.zip" % i

            f1 = Key(b,k)
            f1.set_contents_from_filename(
                '/data/MAgE_20150331_023936.log'
            );

    def logs_on_s3(self):

       s3 = boto.connect_s3()
       b = s3.get_bucket('dmp-log-analysis')

       return len([k for k in b.list(self.customer)])
    

