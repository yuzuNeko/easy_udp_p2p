import socket


udp=socket.socket(socket.AF_INET,socket.SOCK_DGRAM)
udp.bind(('0.0.0.0',9100))
peers=[]
data=''
print('udp服务器开始启动')
while 1:
	data,addr=udp.recvfrom(1024)
	print(data,addr)
	peers.append(addr)
	if len(peers)==2:
		udp.sendto(str(peers[0]).encode('utf-8'),peers[1])
		udp.sendto(str(peers[1]).encode('utf-8'),peers[0])
		print('tell address complete!')
		break