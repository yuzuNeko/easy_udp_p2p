import socket


#首先不写bind的形式
s=socket.socket(socket.AF_INET,socket.SOCK_DGRAM)
server_domain=''
remote_server=(server_domain,9100)
s.sendto(b'hello,server',remote_server)
data,addr=s.recvfrom(1024)

data=data.decode('utf-8')
addr=''
exec('addr={}'.format(data))
print(addr)
s.sendto(b'shake hand',addr)
from threading import Thread 
import time
def send_msg(sock,addr):
	while 1:
		time.sleep(10)
		sock.sendto(b'news',addr)
t=Thread(target=send_msg,args=(s,addr))
t.start()
while 1:
	data,addr=s.recvfrom(1024)
	print(data,addr)
	time.sleep(10)



