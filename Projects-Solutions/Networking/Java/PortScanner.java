/*Port Scanner
* Enter an IP address and a port range where the program will then attempt to 
* find open ports on the given computer by connecting to each of them. 
* On any successful connections mark the port as open.
*/

import java.net.*;
import java.io.*;

public class PortScanner {
	//try to connect to port and if connection is successful return true
	public static boolean connectToPort(String hostname, int port){
		try{
			System.out.println("Trying port " + port + " on " + hostname);
			Socket client = new Socket(hostname, port);
			System.out.println("Port " + port + " on " + hostname + " is open");
			client.close();
			return true;
		}catch(UnknownHostException e){
			System.out.println(e.getMessage());
		}catch(IOException e){
			System.out.println(e.getMessage());
		}
		return false;
	}
	
	//display open ports
	public static void scanPorts(String hostname, int startPort, int endPort){
		int[] allOpenPorts = new int[endPort - startPort];
		int index = 0;
		for(int port = startPort; port <= endPort; port++){
			if(connectToPort(hostname, port)){
				allOpenPorts[index++] = port;
			}
		}
		
		clearScreen();
		System.out.println("Open ports");
		for(int i = 0; i < index; i++){
			System.out.println(allOpenPorts[i]);
		}
	}
	
	//clear screen
	public static void clearScreen(){
    	System.out.printf("\033[2J");
  		System.out.printf("\033[%d;%dH",0,0);
	}


	public static void main(String[] args){
		if(args.length < 3){
			System.out.println("Usage: java PortScanner <ip address> <start port> <end port>");
			System.out.println("java PortScanner 127.0.0.1 1 10000");
		}else{
			String ipAddress = args[0];
			final int startPort = Integer.parseInt(args[1]);
			final int endPort = Integer.parseInt(args[2]);
			scanPorts(ipAddress,startPort,endPort);
		}
	}
}
