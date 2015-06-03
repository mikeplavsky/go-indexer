from ExecuteQuery import Execute
from date_utils import DateUtils
import fs

global run_count
run_count = 0

class InitialState:
    def __init__(self):
        import os,subprocess
        compact = os.path.abspath(os.path.join(fs.curdir(),"../compact_mongo.sh"))
        subprocess.call(["sudo",compact])
        
        
    def DriveFreeBytes(self):
        return fs.get_free_bytes()
    
    def DbSize(self):
        return fs.get_db_size()

class PerformanceCounters(Execute, DateUtils):
    def __init__(self, begin_date_total, begin_drive_free_bytes, begin_db_size):
        self.begin_date_total = begin_date_total
        self.begin_drive_free_bytes = int(begin_drive_free_bytes)
        self.begin_db_size = int(begin_db_size)
        
    def get_dataset(self):
        keys = ["Id","TimeRunTotal", "TimeRunCurrent", "DiskUsed", "DBSizeIncreased"]
        global run_count
        run_count += 1
        
        return [{
            "Id":run_count,
            "TimeRunTotal":self.now() - self.express_date(self.begin_date_total),
            "TimeRunCurrent":self.now() - self.express_date(self.begin_date_total),
            "DiskUsed":fs.sizeof_fmt(self.begin_drive_free_bytes - fs.get_free_bytes()),
            "DBSizeIncreased": fs.sizeof_fmt( fs.get_db_size() - self.begin_db_size )
        }], keys
        
    