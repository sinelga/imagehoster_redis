GOPATH=$GOPATH:/home/juno/git/imagehoster_redis go test -v



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


bin/contents_feeder_redis --site=www.teinit.info --deltahours=50
bin/contents_feeder_redis --site=teinit.info --deltahours=50

bin/sitemap_maker -site www.teinit.info -limit 343 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_www.teinit.info.info.xml
bin/sitemap_maker -site teinit.info -limit 343 >/home/juno/git/8_fi_FIporno_desk/dist/sitemap_teinit.info.xml

#common38ang
bin/contents_feeder_redis --site=www.panee.fi --deltahours=50
bin/contents_feeder_redis --site=panee.fi --deltahours=50

bin/contents_feeder_redis --site=www.suomalaisia.com --deltahours=50
bin/contents_feeder_redis --site=suomalaisia.com --deltahours=50








