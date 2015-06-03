import os,json,re
from date_utils import DateUtils

class LogChecker(DateUtils):
    def startLogCheck(self):
        global begin_date
        begin_date = self.now()
        return begin_date
        
    def curdir(self):
        return os.path.dirname(os.path.abspath(__file__))
        
    def process_empty(self, ret):
        if len(ret):
            return ret
        else:
            return "nothing"

    def searchEsInFile(self, error, path):
        global begin_date
        errors = ""
        with open(path) as f:
            for line in  f:
                dt = re.search("^\[([\d]{4}-[\d]{2}-[\d]{2} [\d]{2}:[\d]{2}:[\d]{2},[\d]{3})\]", line)
                if dt:
                    cur_dt = self.es_date(dt.group(1))
                    if (cur_dt > begin_date) and re.search(error, line):
                        errors += line
                    
        return self.process_empty(errors)
    
    def searchExpressMoreThanInFile(self,code_key, code_value, path):
        global begin_date
        errors = ""
        with open(os.path.join(self.curdir(), "../../%s" % path)) as f:
            for line in f:
                data = json.loads(line)
                if self.express_date(data["timestamp"]) > begin_date and data["res"][code_key] > int(code_value):
                    errors += "Error %d for rqquest: %s\n" % (data["res"][code_key], data["req"]["originalUrl"])
        
        return self.process_empty(errors)
    
    def searchMongodbInFile(self, error, path):
        global begin_date
        errors = ""
        with open(path) as f:
            for line in  f:
                if re.search("^[A-Z][a-z]{2} [A-Z][a-z]{2} [\d]{2} [\d]{2}:[\d]{2}:[\d]{2}.[\d]{3} \[[a-z]*\] ", line):
                    cur_dt = self.mongo_date(line[0:line.find('.')])
                    if (cur_dt > begin_date) and re.search(error, line):
                        errors += line
        return self.process_empty(errors)
    
    def jaydata(self,path):
        return "nothing"
        