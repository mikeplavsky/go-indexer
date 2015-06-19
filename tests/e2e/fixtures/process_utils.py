import subprocess, os

class EnvironmentVariablesValues:

    def table(self, rows):

        header = rows[0]

        for h in header:

            if h.rfind("?") + 1 == len(h):

                env = h.rstrip("?")
                if (os.environ.has_key(env)):
                    setattr(self, env, lambda: os.environ[env])
            else:
                setattr(self, "set%s" % h,  lambda x, h=h: os.environ.update({h:x}))

class ProcessUtils:

    def runCommand( self, cmd ):

        subprocess.call( [ cmd ] )

    def runCommandWithArgs(self, cmd, args):

        subprocess.call([cmd] + args.split(","))

    def runCommandEx( self, args ):

       return subprocess.check_output( args )
