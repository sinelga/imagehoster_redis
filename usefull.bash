vm.overcommit_memory=1 in /etc/sysctl.conf !!!

GOPATH=$GOPATH:/home/juno/git/imagehoster_redis go test -v


5 6 * * * cd /home/juno/git/imagehoster_redis && bin/sitemap_updater



 find . -type f |grep '^./0' |awk '{print "mv "$0" dest"$0  }' 
 find . -type f |grep '^./0' |awk '{print "mkdir -p `dirname "$0"`; mv "$0" dest"$0  }'
 
 :1,96s/dest.\/0/.\//
1,96s/dirname .\/0/dirname .\//

find . -type d -empty -delete


scp 10.132.97.23:git/imagehoster/weber_production.db .
scp 10.132.97.23:git/imagehoster/upload.tar.gz .

tar zxfv  upload.tar.gz
rm upload.tar.gz

SQL
use mysql
drop database weber_production;
create database weber_production;
use weber_production;

update characters set moto='',description ='',region_id=0,adv_phone_id=0 where Topic='sex' and sex='female';

mysqldump --order-by-primary --compress --single-transaction --skip-triggers -uasterisk -p weber_production >weber_production.db

cat weber_production.db |mysql -uasterisk -p weber_production

rm weber_production.db



bin/contents_feeder_redis --site=www.alastonkuvia.info --deltahours=50
bin/contents_feeder_redis --site=alastonkuvia.info --deltahours=50

bin/sitemap_maker -site www.alastonkuvia.info -limit 333 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_www.alastonkuvia.info.xml
bin/sitemap_maker -site alastonkuvia.info -limit 333 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_alastonkuvia.info.xml


bin/contents_feeder_redis --site=www.kotirouvat.com --deltahours=50
bin/contents_feeder_redis --site=kotirouvat.com --deltahours=50

bin/sitemap_maker -site www.kotirouvat.com -limit 343 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_www.kotirouvat.com.info.xml
bin/sitemap_maker -site kotirouvat.com -limit 343 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_kotirouvat.com.xml

#common38ang
bin/contents_feeder_redis --site=www.panee.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=panee.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.panee.fi --deltahours=50 #remake

bin/contents_feeder_redis --site=www.suomalaisia.com --deltahours=50 #remake
bin/contents_feeder_redis --site=suomalaisia.com --deltahours=50 #remake

bin/contents_feeder_redis --site=www.alasti.eu --deltahours=50 #remake
bin/contents_feeder_redis --site=alasti.eu --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.alasti.eu --deltahours=50 #remake

bin/contents_feeder_redis --site=www.nuori.co --deltahours=50 #remake
bin/contents_feeder_redis --site=nuori.co --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.nuori.co --deltahours=50 

bin/contents_feeder_redis --site=www.allfreexxx.com --deltahours=50 #remake
bin/contents_feeder_redis --site=allfreexxx.com --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.allfreexxx.com --deltahours=50 #remake


bin/contents_feeder_redis --site=www.klickhere.mobi --deltahours=50 #remake
bin/contents_feeder_redis --site=klickhere.mobi --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.klickhere.mobi --deltahours=50 #remake

bin/contents_feeder_redis --site=www.bestsex.cc --deltahours=50 #remake
bin/contents_feeder_redis --site=bestsex.cc --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.bestsex.cc --deltahours=50 #remake

bin/contents_feeder_redis --site=www.rouvat.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=rouvat.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.rouvat.fi --deltahours=50 #remake

bin/contents_feeder_redis --site=www.vittuja.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=vittuja.fi --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.vittuja.fi --deltahours=50 #remake

bin/contents_feeder_redis --site=www.mostpopularsites.org --deltahours=50 #remake
bin/contents_feeder_redis --site=mostpopularsites.org --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.mostpopularsites.org --deltahours=50 #remake

#common39ang
bin/contents_feeder_redis --site=www.pornocopia.net --deltahours=50 #remake
bin/contents_feeder_redis --site=pornocopia.net --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.pornocopia.net --deltahours=50 #remake


bin/contents_feeder_redis --site=www.redtubepornos.org --deltahours=50 #remake
bin/contents_feeder_redis --site=redtubepornos.org --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.redtubepornos.org --deltahours=50 #remake

bin/contents_feeder_redis --site=www.xporno18.com --deltahours=50 #remake
bin/contents_feeder_redis --site=xporno18.com --deltahours=50 #remake
bin/contents_feeder_redis --site=chat.xporno18.com --deltahours=50 #remake


bin/contents_feeder_redis --site=www.sexualwish.com --deltahours=50 #remake
bin/contents_feeder_redis --site=sexualwish.com --deltahours=50  #remake
bin/contents_feeder_redis --site=chat.sexualwish.com --deltahours=50 #remake

bin/contents_feeder_redis --site=www.seksideitti.fi --deltahours=50
bin/contents_feeder_redis --site=seksideitti.fi --deltahours=50 
bin/contents_feeder_redis --site=chat.seksideitti.fi --deltahours=50

#common40ang
bin/contents_feeder_redis --site=teinit.info --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.teinit.info --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.teinit.info --deltahours=50
bin/contents_feeder_redis --site=kotirouvat.com --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.kotirouvat.com --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.kotirouvat.com --deltahours=50
bin/contents_feeder_redis --site=alastonkuvia.info --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.alastonkuvia.info --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.alastonkuvia.info --deltahours=50
bin/contents_feeder_redis --site=sinkkutreffit.org --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.sinkkutreffit.org --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.sinkkutreffit.org --deltahours=50 
bin/contents_feeder_redis --site=alaston.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.alaston.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.alaston.fi --deltahours=50
bin/contents_feeder_redis --site=pimppi.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.pimppi.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.pimppi.fi --deltahours=50
bin/contents_feeder_redis --site=alastomat.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.alastomat.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.alastomat.fi --deltahours=50
bin/contents_feeder_redis --site=seksivideo.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.seksivideo.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.seksivideo.fi --deltahours=50
bin/contents_feeder_redis --site=runkkaus.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.runkkaus.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.runkkaus.fi --deltahours=50
bin/contents_feeder_redis --site=karvainen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.karvainen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.karvainen.fi --deltahours=50
bin/contents_feeder_redis --site=munaa.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.munaa.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.munaa.fi --deltahours=50
bin/contents_feeder_redis --site=teijan.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.teijan.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.teijan.fi --deltahours=50
bin/contents_feeder_redis --site=naisen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.naisen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.naisen.fi --deltahours=50
bin/contents_feeder_redis --site=vittua.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.vittua.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.vittua.fi --deltahours=50
bin/contents_feeder_redis --site=antaa.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.antaa.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.antaa.fi --deltahours=50
bin/contents_feeder_redis --site=nussi.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.nussi.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.nussi.fi --deltahours=50



#common41ang
bin/contents_feeder_redis --site=pippeli.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.pippeli.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.pippeli.fi --deltahours=50
bin/contents_feeder_redis --site=seksitarina.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.seksitarina.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.seksitarina.fi --deltahours=50bin/contents_feeder_redis --site=kiima.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.kiima.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.kiima.fi --deltahours=50
bin/contents_feeder_redis --site=miehen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.miehen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.miehen.fi --deltahours=50
bin/contents_feeder_redis --site=peniksen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.peniksen.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.peniksen.fi --deltahours=50
bin/contents_feeder_redis --site=seksiverkko.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.seksiverkko.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.seksiverkko.fi --deltahours=50





bin/contents_feeder_redis --site=sub.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=www.sub.fi --deltahours=50 && sleep 90 && bin/contents_feeder_redis --site=chat.sub.fi --deltahours=50
