import boto
from boto.s3.key import Key

def upload_logs_to_S3():
    
    keys = [
        "CustomerA/MachineB/Agent%s/MAgE_20150331_023936.zip" % i 
        for i in range(0,10)
    ]

    s3 = boto.connect_s3()
    b = s3.get_bucket('dmp-log-analysis')

    [b.delete_key(k) for k in keys]   

    for k in keys:

        f1 = Key(b,k)
        f1.set_contents_from_filename(
            'MAgE_20150331_023936.log'
        );
