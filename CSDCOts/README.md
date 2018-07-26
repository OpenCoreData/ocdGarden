



Ref: https://hub.docker.com/r/nawer/blazegraph/

docker stack deploy -c stack.yml nawer/blazegraph 
or 
docker-compose -f stack.yml up
Wait for it to initialize completely, and visit 
http://swarm-ip:9999, 
http://localhost:9999, 
or 
http://host-ip:9999 

(as appropriate)



SPARQL updates

PREFIX csdco: <http://csdco.org/>
INSERT DATA
{ 
 <http://csdco.org/res/14>  <http://csdco.org/voc#pred1>  <http://csdco.org/res/17> . 
}


PREFIX csdco: <http://csdco.org/>
DELETE DATA
{
 <http://csdco.org/res/14>  <http://csdco.org/voc#pred1>  <http://csdco.org/res/17> . 
}
