#!/bin/bash
if [ "$EUID" -ne 0 ]
then
    echo "Please run this script with sudo."
    exit 1
fi

# Prompt the user for a password
read -sp 'Enter a password for the MariaDB root user: ' password
echo



# Install MariaDB
if ! command -v mariadb &> /dev/null
then
    echo "MariaDB is not installed. Installing now..."
    sudo apt-get update
    sudo DEBIAN_FRONTEND=noninteractive apt-get install -y mariadb-server
    echo "MariaDB installed successfully."
    echo "MariaDB is already installed."

# Configure MariaDB
echo "Configuring MariaDB..."
sudo mysql_secure_installation <<EOF

$password
n
Y
Y
Y
Y
EOF
echo "MariaDB configured successfully."

else
    echo "MariaDB is already installed."
fi 
    # Check if the user wants to reset the root password
    read -p "Do you want to reset the MariaDB root password, this will restart your mariaDB as well ? [y/n]: " reset_password
    if [ "$reset_password" == "y" ]
    then
 systemctl stop mariadb
 systemctl set-environment MYSQLD_OPTS="--skip-grant-tables --skip-networking"
 systemctl start mariadb
 mysql -u root --execute="FLUSH PRIVILEGES; UPDATE mysql.user SET password = PASSWORD('${password}') WHERE user = 'root'; \
UPDATE mysql.user SET authentication_string = '' WHERE user = 'root';
UPDATE mysql.user SET plugin = '' WHERE user = 'root';"
systemctl unset-environment MYSQLD_OPTS
systemctl restart mariadb
fi 
### final test to see if everything went well 
echo "now checking if mariadb works with newpasswd "

if [[ $(mysql -u root -p${password} --execute="exit" 2>&1) == *"ERROR"* ]]; then
  echo "MySQL password incorrect or command execution failed."
else
  echo "MySQL password correct and command executed successfully."
fi
echo "your password is " ${password}