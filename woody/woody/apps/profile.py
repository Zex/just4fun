from pyspark.context import SparkContext
from pyspark.conf import SparkConf
from pyspark.streaming.context import StreamingContext
from pyspark.streaming.kafka import KafkaUtils
from woody.common.config import Config
from py4j.protocol import Py4JJavaError
from datetime import datetime
from random import randint

class Profile(object):
    """Profile aggregates
    """
    def __init__(self):
        conf = SparkConf()
        conf.set("spark.scheduler.mode", Config.spark_sched_mode)
        conf.set("spark.scheduler.pool", Config.spark_sched_pool)
        conf.set("spark.scheduler.allocation.file", Config.spark_sched_file)
        self._sc = SparkContext(appName=self.__class__.__name__+str(randint(10,20)), conf=conf)
        self._sc.setLogLevel('WARN')
        self._ssc = StreamingContext(self._sc, Config.ssc_duration)
        self._from_kfk()

    def _from_kfk(self):
        kfk_params = {
            'metadata.broker.list': Config.kafka_brokers,
            'group.id': self.__class__.__name__,
            }
        self._ds = KafkaUtils.createDirectStream(self._ssc, ["user_imsg"], kfk_params)
        fmt = "%Y-%m-%d %H:%M:%S"
        #dates = self._ds.flatMap(lambda x: [datetime.strptime(
        #    ''.join(x[1].split('.')[:1]), fmt)])
        # <datetime>: "checked in"
        dates = self._ds.map(lambda x: datetime.strptime(
            ''.join(x[0].split('.')[:1]), fmt))
        dates = dates.map(lambda x: (x, 1)).reduceByKey(lambda x, y: y+x)
        dates.pprint()
        #dates.saveAsTextFiles(self.__class__.__name__)

    def foreach_rdd(self, time, rdd):
        print(time, rdd)

    def run(self):
        try:
            self._ssc.start()
            self._ssc.awaitTermination()
        except Py4JJavaError as e:
            print("start failed", e)
        except KeyboardInterrupt:
            print("shutdown")
        #    self._ssc.stop()
        except Exception as e:
            print("unhandled: ", e)

if __name__ == '__main__':
    app = Profile()
    app.run()