#### Globalsync ( pronounced globalsynch)

----
an opensource software Developed by Arash Mohammadreza Mohammadi for managing Data could be file or folder  across a cluster of servers, so that for example you may need to have a set of data that shall be synchronized across sets of servers this software will let you do it .


1. it is capable to find the closest available source that match the requirements 
2. it will offer the Destinations that has capability to recieve that data based on the available capacity 
3. the software is agentless can communicates  with servers through ssh and leverage the rsync to perform data transfer
----
##### to start 
1. you need to have latest version of go install 1.20+ if you wish to build it yourself , otherwise the the binaries are available  in bin folder
2. this software uses mysql/mariadb  10.X, tables schemas are inside the ./db , it may generates some  Errors  during mock data entry, because of the auto increment feature for the id, on tables,  one may find a solution for it by either manually checking depend tables and matching foreign keys with existence one and matching them or programmatically  through scripting of this part 
run the main.go , one may wish to create a systemd runner out of binary file or even create a k8s manifest and docker file these things are still under dev will be published soon 

### todo

1. create a database, test data creation through scripts or inside the program 
2. create docker image and kubernetes manifest for the ease of deployment 
3. create a webhook for copy action based on data sanitization on RDBMS and user input validator 
4. create a frontend for it 

##### support
this software has been launched as beta there is no support, and its provided on  basis of "AS IS" without support or warranty  and all risk taken to use, redistribute is on the user side and developer has no responsibility regarding use, deployment, distribution in either test or productions environment.

## license 
this software MIT
