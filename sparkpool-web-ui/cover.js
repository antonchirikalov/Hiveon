module.exports = {
  kafka: {
    worker_topic: 'hiveos_eth_stats_prod',
    exchange_topic: 'hiveos_eth_exchange_prod',
  },
  influx: {
    host: 'influx',
    port: 8086,
    database: 'minerdash',
    username: 'admin',
    password: 'UadZOn2Hsy1Dyt07',
  },
  redis: {
    host: 'redis',
    port: 6379,
    password: 'kAloC73uB0C7puSx',
    db: 1,
  },
  sequelize: { // 只读
    dialect: 'mysql', // support: mysql, mariadb, postgres, mssql
    database: 'oauth',
    host: 'rm-2ze41sh9k2oquw8c4o.mysql.rds.aliyuncs.com',
    port: '3306',
    username: 'oauth_rw',
    password: '%0ODaSZVPr%Yq1DG',
    timezone: '+08:00',
    charset: 'utf8',
    pool: {
      max: 20,
      min: 5,
      idle: 30000,
      acquire: 20000,
    },
  },
  sequelize2: { // 只读
    dialect: 'mysql', // support: mysql, mariadb, postgres, mssql
    database: 'hiveos_eth',
    host: 'hiveos-eth-prod.cyzben5dhs1h.eu-central-1.rds.amazonaws.com',
    port: '3306',
    username: 'hiveos_eth_ro',
    password: '7C9RAXMGQKiOKZ1A',
    timezone: '+08:00',
    charset: 'utf8',
    pool: {
      max: 200,
      min: 25,
      idle: 30000,
      acquire: 20000,
    },
  },
  sequelize3: { // 只读
      dialect: 'mysql', // support: mysql, mariadb, postgres, mssql
      database: 'block',
      host: 'rm-2zefa4845bwys9fb9io.mysql.rds.aliyuncs.com',
      port: '3306',
      username: 'ethscan_read',
      password: 'OLyEwNnx75uWtGkI',
      timezone: '+08:00',
      charset: 'utf8',
      pool: {
        max: 20,
        min: 5,
        idle: 30000,
        acquire: 20000,
      },
    },
};
