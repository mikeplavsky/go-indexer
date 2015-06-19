from process_utils import ProcessUtils
from os import path
import json
import sys

crypt_script = path.join( path.dirname( __file__ ), '..', '..', 'crypt', 'crypt-rsa.js' )

if sys.platform == 'win32':

  node_exe = path.join( path.dirname( __file__ ), '..', '..', '..', 'vendor', 'nodejs', 'node.exe' )
else:
  node_exe = 'node'


class CryptUtils:

  public_key = None
  private_key = None

  def GenerateKeys( self, bits ):

    result = ProcessUtils().runCommandEx( [ node_exe, crypt_script, 'genkeys', str(bits) ] )

    keys = json.loads( result )[ 'keys' ]
    CryptUtils.public_key = keys[ 0 ]
    CryptUtils.private_key = keys[ 1 ]

    return CryptUtils.public_key

  def GetPublicKey( self ):

    return CryptUtils.public_key

  def GetPrivateKey( self ):

    return CryptUtils.private_key

  def Encrypt( self, public_key, data ):

    result = ProcessUtils().runCommandEx( [ node_exe, crypt_script, 'encrypt', public_key, data ] )
    return json.loads( result )[ 'data' ]

  def Decrypt( self, private_key, data ):

    result = ProcessUtils().runCommandEx( [ node_exe, crypt_script, 'decrypt', private_key, data ] )
    return json.loads( result )[ 'data' ]
