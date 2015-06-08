from waferslim.converters import converter_for
import boto

class S3Buckets(object):
    
    def query(self):

        s3 = boto.connect_s3()

        bs = s3.get_all_buckets()
        res = [[['Name',x.name]] for x in bs]

        converter = converter_for(list)
        return converter.to_string(res)

