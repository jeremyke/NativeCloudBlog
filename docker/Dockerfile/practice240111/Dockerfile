FROM elasticsearch:6.4.2
LABEL maintainer="maxwellke <jeremyke@qq.com>" \
    description="elasticsearch6-4-2 with ik"
ADD https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.4.2/elasticsearch-analysis-ik-6.4.2.zip ./
RUN yum install -y unzip
RUN unzip ./elasticsearch-analysis-ik-6.4.2.zip -d ./plugins/elasticsearch-analysis-ik
RUN rm -rf ./elasticsearch-analysis-ik-6.4.2.zip