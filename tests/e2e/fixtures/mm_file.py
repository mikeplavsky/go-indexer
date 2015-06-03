import mmap
from tempfile import TemporaryFile

class MMFile:

    tempFile=None

    def mmwrite(self, data):

        self.tempFile.write(data)

    def mmflush(self):

        self.tempFile.flush()

    def mmseek(self, pos):

        self.tempFile.seek(pos)

    def mmgetBody(self):

        self.mmapped_file_as_string = mmap.mmap(self.tempFile.fileno(), 0, access=mmap.ACCESS_READ)
        return self.mmapped_file_as_string

    def mmclose(self):

        self.mmapped_file_as_string.close()
        self.tempFile.close()

    def mmcreateFile(self):

        t = TemporaryFile()
        self.tempFile = t
        return t
