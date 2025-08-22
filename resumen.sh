####################################################bashrc##################
case $- in
    *i*) ;;
      *) return;;
esac
HISTCONTROL=ignoreboth
shopt -s histappend
# for setting history length see HISTSIZE and HISTFILESIZE in bash(1)
HISTSIZE=1000
HISTFILESIZE=2000

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# If set, the pattern "**" used in a pathname expansion context will
# match all files and zero or more directories and subdirectories.
#shopt -s globstar

# make less more friendly for non-text input files, see lesspipe(1)
#[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "${debian_chroot:-}" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
    xterm-color|*-256color) color_prompt=yes;;
esac

# uncomment for a colored prompt, if the terminal has the capability; turned
# off by default to not distract the user: the focus in a terminal window
# should be on the output of commands, not on the prompt
#force_color_prompt=yes

if [ -n "$force_color_prompt" ]; then
    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
	# We have color support; assume it's compliant with Ecma-48
	# (ISO/IEC-6429). (Lack of fsuch support is extremely rare, and such
	# a case would tend to support setf rather than setaf.)
	color_prompt=yes
    else
	color_prompt=
    fi
fi

if [ "$color_prompt" = yes ]; then
    # PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
    PS1='[\w]\n\[\e[94;1m\]\A\[\e[0m\]\$'
else
    PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
unset color_prompt force_color_prompt

# If this is an xterm set the title to user@host:dir
case "$TERM" in
xterm*|rxvt*)
    PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
    ;;
*)
    ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
fi
# some more ls aliases
alias u='sudo apt update&&sudo apt upgrade'
alias l='ls -Cfa'
alias s='ssh oscar@192.168.1.126'
alias c='clear'
alias b='sudo vim .bashrc'
alias h='history|less'
alias hvim='vim .bash_history'
alias g='go run main.go'
alias p='sudo poweroff'
alias r='sudo reboot'
alias n='cd /var/www/html&&netbeans'
alias db='sudo -u postgres psql'
alias backup='mysqldump -uoscar -pemi loft > /media/oscar/OSCAR/OSCAR/loft0001.sql&&cd /media/oscar/OSCAR/OSCAR/&&ls -l loft*'
alias backupC='mysqldump -uoscar -pemi loft > /media/oscar/CB07-D711/loft0001.sql&&cd /media/oscar/CB07-D711/&&ls -l loft*'
alias backupPC='mysqldump -uoscar -pemi loft > Documentos/loft0001.sql&&cd /Documentos/&&ls -l loft*'
alias v='cd /etc/vim&&sudo vim vimrc'

if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

################################bash_history########################################
#instalacion de debian 13 trixie
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

sudo apt install snapd
sudo snap install pgadmin4
#cuando instalo el pgadmin creo primero usuario como puse arriba
#sudo -i -u postgres despues pongo createuser --interactive;
#con git es otra cosa
#usuario angorita password mengeche2024 
#despues en consola coloco lo que sigue
git clone https://github.com/angorita/loft.git

git status
git add .
git commit -m 'Primer Commit'
git push
touch .gitignore
git clone http://github.com/angorita/godesde0.git