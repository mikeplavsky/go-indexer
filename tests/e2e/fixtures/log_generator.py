import os
from tempfile import NamedTemporaryFile

log = None

class Log:
    def create(self):
        global log
        log = NamedTemporaryFile(delete=False);
        return log.name
    
    def close(self):
        global log
        log.close()

class Generate:

    def add_field(self, str, value):
        return str + value + "\t"
    
    def table(self, rows):
        
        header = rows[0]
        for h in header:
            setattr(self, "set%s" % unicode.replace(h,h[0],h[0].upper(),1),  lambda x:x)

        for row in rows[1:]:

            log_record = ""

            for idx in range(len(header)):
                coll_name = header[idx]
                print coll_name
                print row[idx]
                log_record = self.add_field(log_record, row[idx])
        
            global log
            
            log.write(log_record+"\n")
        log.flush();
