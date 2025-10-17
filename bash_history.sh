##############################################
#instalacion de debian 13 trixie
##############################################
#se instala igual pero ...en la parte de 
#solo clickeo aplicaciones estandar, nada mas
#ver resto en cuaderno de rocio con el pajaro...
sudo apt -y install gnome-core gdm3

sudo apt install gdebi gdebi-core synaptic
sudo nano /etc/default/grub
#editamos lo siguiente 
#cambio timeout a 0
#quiet splash
#GFXMODE 1920X1080 
#actualizo despues con update-grub
sudo update-grub
sudo apt autoremove
sudo apt install btop
btop
sudo apt install apache2
systemctl start apache2
systemctl status apache2.service 
sudo apt install postgresql postgresql-contrib
#con esto puedo buscar bien
############################ POSTGRESQL ################################
apt search postgresql
apt search postgresql|grep postgresql
sudo systemctl status postgresql
sudo systemctl stop postgresql
#edito el archivo de configuracion 
sudo nano /etc/postgresql/17/main/postgresql.conf 
# aca se pone listen_addresses ='localhost'

sudo -i -u postgres
#aca entramos en postgres para crear un usuario.
#createuser --interactive pregunta y colocamos lo que corresponde
#createdb nombrebasedatos
#alter user nombreusuario with password 'contrasenia';
#ALTER ROLE 
#si quiero crear privilegios sobre una base de datos
#grant all privileges on database "nombrebasedatos" to oscar;


sudo systemctl status postgresql
sudo systemctl start postgresql
#posteriormente hay que crear un usuario
# 1. Actualiza los índices de paquetes
sudo apt update

# 2. Instala el módulo de Apache para PHP y la extensión MySQL/MariaDB
# Usa el número de versión (ej. php8.2) si lo tienes instalado.
# Si no, usa 'php' para instalar la versión por defecto de tu distribución.
sudo apt install libapache2-mod-php php-mysql
sudo apt install snapd
sudo snap install pgadmin4
#cuando instalo el pgadmin creo primero usuario como puse arriba
#sudo -i -u postgres despues pongo createuser --interactive;
#con git es otra cosa
#usuario angorita password mengeche2024 
#despues en consola coloco lo que sigue
git clone https://github.com/angorita/loft.git

bash subir.sh 
############################### SPOTIFY ###################################
sudo snap install spotify
sudo apt-get update && sudo apt-get install spotify-client
go mod init SSC
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3
air
ls
cd ..;\
cd ..
c
mdb-tables NeoSSC.mdb 
mdb-export NeoSSC.mdb PARTOS >partos.csv
cat partos.csv 
c
echo -e ".separator ,\n.import partos.csv partos"|sqlite3 partos.db
l
################### MDB TOOL'S ########################
sudo apt install mdbtools
mdbtools
mbdtools
cd Descargas/
mdb-tables NeoSSC.mdb 
mdb-export NeoSSC.mdb NEONATOLOGO >partos.csv 
mdb-export NeoSSC.mdb NEONATOLOGO >neonatologo.csv
cat neonatologo.csv 
echo -e".separator ,\n.import neonatologo.csv neonatologo|sqlite3 partos.db
;
echo -e ".separator ,\n.import neonatologo.csv neonatologo"|sqlite3 partos.db
h
c
mdb-tables NeoSSC.mdb 
mdb-export NeoSSC.mdb OSOCIAL OBS >osocial.csv obs.csv
c
mdb-export NeoSSC.mdb OSOCIAL >osocial.csv 
mdb-export NeoSSC.mdb OBS >obs.csv
ls -la

################### PHP MY ADMIN ######################
# 1. Actualiza los índices de paquetes
sudo apt update
# 2. Instala el módulo de Apache para PHP y la extensión MySQL/MariaDB
# Usa el número de versión (ej. php8.2) si lo tienes instalado.
# Si no, usa 'php' para instalar la versión por defecto de tu distribución.
sudo apt install libapache2-mod-php php-mysql
