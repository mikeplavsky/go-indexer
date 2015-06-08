from waferslim.converters import converter_for

class S3Buckets(object):
    
    def query(self):

        converter = converter_for(list)
        return converter.to_string(
               [[['Name','One'],['Last','Two']],    
                [['Name','Two'],['Last','Three']]])    

