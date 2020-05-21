# Crawlers de Finanças Pessoais

Crawlers de serviços financeiros para utilização pessoal.

## Objetivo

O objetivo desse repositório é juntar scripts que possam automatizar a coleta de dados e informações sobre transações financeiras pessoais.

Cada pacote referente a diferentes serviços financeiros realiza a mineiração de dados de formas diferentes, exportando dados comumente no formato JSON.

### Nubank

O pacote do Nubank utiliza os comandos do _android debug bridge_ (`adb shell`) para se conectar ao app do nubank em um aparelho android e realizar o scraping nas telas de gastos.

Os dados obtidos são exportados na linha de comando em formato JSON.

### Clear Corretora

O pacote da Clear utiliza o Godog e Selenium Webdriver para preencher dados na web e pegar registros de transações financeiras de vários períodos operados.

Os dados de transações financeiras são guardados em formato JSON, podendo ser convertido para xls através de pacotes helpers.
