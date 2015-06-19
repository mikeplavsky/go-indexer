import os, platform,ctypes

def get_free_bytes(folder = "/"):
    if platform.system() == 'Windows':
        free_bytes = ctypes.c_ulonglong(0)
        ctypes.windll.kernel32.GetDiskFreeSpaceExW(ctypes.c_wchar_p(folder), None, None, ctypes.pointer(free_bytes))
        return free_bytes.value
    else:
        st = os.statvfs(folder)
        return st.f_bavail * st.f_frsize
        
def get_folder_size(folder = "/"):
    folder_size = 0
    for (path, dirs, files) in os.walk(folder):
      for file in files:
        filename = os.path.join(path, file)
        folder_size += os.path.getsize(filename)
    return folder_size
        
def sizeof_fmt(num):
    for x in ['bytes','KB','MB','GB','TB']:
        if num < 1024.0:
            return "%3.1f %s" % (num, x)
        num /= 1024.0        
        
def get_db_size():
    db_folder = "/var/lib/mongodb"
    if platform.system() == 'Windows':
        cwd = os.path.dirname(__file__)
        db_folder = os.path.join(cwd,"../../mongo/data") #not tested on Win!!!
    return get_folder_size(db_folder)
 
def curdir():
    return os.path.dirname(os.path.abspath(__file__))
