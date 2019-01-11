FROM zlynch/spark:2.4.0

ENV WOODY_ROOT /opt/woody
WORKDIR $WOODY_ROOT

COPY Makefile requirements.txt $WOODY_ROOT/
COPY woody $WOODY_ROOT/woody 
COPY config $WOODY_ROOT/config
COPY scripts $WOODY_ROOT/scripts

RUN echo "<?xml version=\"1.0\"?> \
<allocations> \
  <pool name=\"production\"> \
    <schedulingMode>FAIR</schedulingMode> \
    <weight>2</weight> \
    <minShare>1</minShare> \
  </pool> \
  <pool name=\"dev\"> \
    <schedulingMode>FAIR</schedulingMode> \
    <weight>2</weight> \
    <minShare>1</minShare> \
  </pool> \
</allocations>" > /opt/spark/conf/fairscheduler.xml

CMD ["scripts/aggr_cli.sh"]