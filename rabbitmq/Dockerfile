FROM rabbitmq:3.13.3-management

RUN apt-get update && \
apt-get install -y curl wget && \
apt-get clean

# Install the RabbitMQ Delayed Message Exchange plugin
RUN curl https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v3.13.0/rabbitmq_delayed_message_exchange-3.13.0.ez > rabbitmq_delayed_message_exchange-3.13.0.ez && \
mv rabbitmq_delayed_message_exchange-3.13.0.ez /opt/rabbitmq/plugins/ 
RUN rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange


