# -*- coding: utf-8 -*-

# © 2014 Dell Inc.
# ALL RIGHTS RESERVED.

import csv
import mmap
import json
from ExecuteQuery import Execute
from http_calls import HttpCall, RestTools
from mm_file import MMFile
from tempfile import TemporaryFile

class CsvAsTable( Execute, HttpCall ):

  def __init__( self, url ):

    self.data = self.GET( self.get_full_url( url ) )

  def get_dataset( self ):

    header = []
    data = []

    reader = csv.reader( self.data.splitlines(), delimiter = ',' )
    for row in reader:

      if header:

        data.append( dict( zip( header, row ) ) )
      else:
        header = row

    return data, header

class CsvImport(RestTools, MMFile):

    def generateAndUploadCsvFile( self, url, template, headers, schema_url, index_name, objectsCount, collectionsCount ):

        fields = self.getObjectStringFields( schema_url, index_name )
        self.generateFile(fields, int(objectsCount), int(collectionsCount))

        req = self.getRequest(url, self.mmgetBody(), json.loads( headers ) if headers else self.http_headers)
        response = self.read(req)

        self.mmclose()

        return response

    def generateFile(self, fields, objectsCount, collectionsCount):

        fieldsCount = len(fields)
        
        if collectionsCount > 0:
            fields.append('collection')

        self.mmcreateFile()

        self.mmwrite('--File-Boundary\r\n')
        self.mmwrite('Content-Disposition: form-data; name="file"; filename="test.txt"\r\n')
        self.mmwrite('Content-Type: text/plain\r\n\r\n')

        self.mmwrite(','.join(fields) + '\r\n')


        for i in range(objectsCount):

            temp = []
            for j in range(fieldsCount):
                temp.append(fields[j] + str(i + 1))
            
            if collectionsCount > 0:
                collectionIndex = i % collectionsCount + 1
                temp.append('collection' + str(collectionIndex))

            self.mmwrite(','.join(temp) + '\r\n')
            self.mmflush()

        self.mmwrite('\r\n--File-Boundary--\r\n')
        self.mmseek(0)
