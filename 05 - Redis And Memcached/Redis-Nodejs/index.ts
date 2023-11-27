// const UPSTASH_REDIS_REST_URL = "https://apn1-enabled-tapir-33281.upstash.io";
// const UPSTASH_REDIS_REST_TOKEN =
//   "AYIBASQgYzkyY2Y5MWUtZTdiMy00NTdiLWExZmUtOThhZTNmOTY0NTc3NDczNTBhODBhN2ZmNGY1ZDljZjdmN2RkZGM5OTRmODE=";

// import { Redis } from "@upstash/redis";
// import https from "https";
// import "isomorphic-fetch";

// const redis = new Redis({
//   url: UPSTASH_REDIS_REST_URL,
//   token: UPSTASH_REDIS_REST_TOKEN,
//   agent: new https.Agent({ keepAlive: true }),
// });
// const p = redis.pipeline();

// p.set("ff", "ff").get("ff");

// async function exec() {
//   const time = Date.now();
//   const res = await p.exec();
//   console.log("elapsed time:", Date.now() - time, "ms");
//   console.log(res);
// }
// exec()
//   .then(() => {})
//   .catch(e => console.log(e));

import pg from "pg";
import fs from "fs";

const config = {
  host: "ap-south-1.026d75ca-22b1-41cd-a4fc-a7bafcc3ee51.aws.ybdb.io",
  port: 5433,
  database: "yugabyte",
  user: "admin",
  password: "ZhNNhr5pMt31oeyvHGML9c2Eu8bb41",
  ssl: {
    rejectUnauthorized: true,
    ca: fs.readFileSync("./root.crt").toString(),
  },
  connectionTimeoutMillis: 5000,
};

let client: any;

async function connect() {
  let time = Date.now();
  console.log(">>>> Connecting to YugabyteDB!");

  try {
    client = new pg.Client(config);

    await client.connect();

    console.log(">>>> Connected to YugabyteDB! in ", Date.now() - time, " ms");
  } catch (err) {
    console.log(err);
  }
}
async function createDatabase() {
  try {
    let time = Date.now();
    let stmt = `CREATE TABLE DemoAccount (
          id int PRIMARY KEY,
          name varchar,
          age int,
          country varchar,
          balance int)`;

    await client.query(stmt);
    console.log(">>>> created DemoAccount table in ", time - Date.now(), " ms");

    stmt = `INSERT INTO DemoAccount VALUES
          (1, 'Jessica', 28, 'USA', 10000),
          (2, 'John', 28, 'Canada', 9000)`;
    time = Date.now();
    await client.query(stmt);

    console.log(">>>> inserted data into DemoAccount table in ", time - Date.now(), " ms");
  } catch (err) {
    console.log(err);
  }
}

async function queryDatabase() {
  try {
    let time = Date.now();
    const res = await client.query(`SELECT * FROM DemoAccount`);
    console.log(">>>> queried DemoAccount table in ", Date.now() - time, " ms");
    console.log(res.rows);
  } catch (err) {
    console.log(err);
  }
}
for (let i = 0; i < 15; i++) {
  connect().catch(e => console.log(e));
}
