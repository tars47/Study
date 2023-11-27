// const PSCALE_DATABASE_URL =
//   'mysql://g5cawnv3ohxteeam1c8c:pscale_pw_W2wqLT5MQJWbG6uwQPslm6fcMb0KWAtfxVT1NsvZtEw@aws.connect.psdb.cloud/db?ssl={"rejectUnauthorized":true}';

// const mysql = require("mysql2/promise");

// let connection;
// async function connect() {
//   try {
//     let time = Date.now();
//     connection = await mysql.createConnection(PSCALE_DATABASE_URL);
//     console.log("time took for connection", Date.now() - time);
//   } catch (e) {
//     console.log(e);
//   }
// }

// async function query() {
//   try {
//     let time = Date.now();
//     let [rows] = await connection.query(`SELECT * FROM users`);
//     console.log(rows);
//     console.log("time took for query", Date.now() - time);
//   } catch (e) {
//     console.log(e);
//   }
// }
//connect().then(() => query().then());

// const { initializeApp } = require("firebase/app");

// const firebaseConfig = {
//   apiKey: "AIzaSyCaOfcjCQYDn54CYZBp40nB8lZKVB9i1GI",
//   authDomain: "beejek-67764.firebaseapp.com",
//   projectId: "beejek-67764",
//   storageBucket: "beejek-67764.appspot.com",
//   messagingSenderId: "498651717107",
//   appId: "1:498651717107:web:3577017fb81a9377a1c686",
//   measurementId: "G-CDXZYCRX69",
// };

// const app = initializeApp(firebaseConfig);

// const firestore = require("firebase/firestore");

// function getSomething() {
//   app
//     .firestore()
//     .collection("merchants")
//     .get()
//     .then(querySnapshot => {
//       querySnapshot.docs.map(doc => {
//         console.log("LOG 1", doc.data());
//         return doc.data();
//       });
//     });
// }
// getSomething();

//col: bookingTypeConfigs
// doc: birthDayParty
//coll: locations
//doc: grapevine (name, address, concurrentBookingsCount, sublevel)
//                  if con-lim is -1
//                 then i have to look at sublevel, else if con-lim is 0 or >
//                then dont look down at sublevel
// coll: rooms
// doc: room1

// Location
// room1
//      rules:{}
// event1: birthday
//         con-lim: 1
// event2: openplay
//         con-lim:40
// event3: firetruck show
//         con-lim:15

// whenTimeslotIsBookedBlockOtherEvents = false;

// rules = [
//   {
//     eventId: "event1",
//     when: "event1 is booked",
//     allow: [],
//     disallow: ["event2", "event3"],
//   },
//   {
//     eventId: "event2",
//     when: "event2 is booked",
//     allow: ["event1"],
//     disallow: ["event3"],
//   },
// ];

// req = {
//   Availability_For: ["Online"],
//   Is_Search_Date_A_Range: false,
//   Org_ID: "jDVCc97TKqLQSyy7C8K4",
//   Config: [
//     { Type: "Location", Path: "/New_Booking_Flow/Locations/Grapevine" },
//     { Type: "Room", Path: "/New_Booking_Flow/Locations/Grapevine/Rooms/001" },
//     {
//       Type: "Booking_Type",
//       Path: "/New_Booking_Flow/Locations/Grapevine/Rooms/001/Booking_Types/Birthday",
//     },
//   ],
//   Search_Dates: ["2023-07-24"],
//   Output_Time_Zone: "America/Chicago",
// };

// let now = new Date().valueOf();
// for (let i = 0; i < 100000000; i++) {}
// let after = new Date().valueOf();
// console.log("nodejs-->", after - now);

// const test = require("./class");

// test.test();

// function generateCombinations(inputArray, currentIndex, currentCombination, outputArray) {
//   if (currentIndex === inputArray.length) {
//     outputArray.push([...currentCombination]);
//   } else {
//     for (const element of inputArray[currentIndex]) {
//       currentCombination.push(element);
//       generateCombinations(inputArray, currentIndex + 1, currentCombination, outputArray);
//       currentCombination.pop();
//     }
//   }
// }

// const inputArray = [[1, 2, 3], [98, 99], ["a", "b", "c"], ["+"], [47, 74]];
// const outputArray = [];

// generateCombinations(inputArray, 0, [], outputArray);

// console.log(outputArray);

class TreeNode {
  constructor(value) {
    this.value = value;
    this.children = [];
  }
}

// Function to get all leg combinations
function getLegCombinations(root) {
  const result = [];

  function dfs(node, currentLeg) {
    if (!node) return;

    currentLeg.push(node.value);

    if (node.children.length === 0) {
      // Leaf node, add the current leg to the result
      result.push([...currentLeg]);
    } else {
      // Non-leaf node, continue the traversal
      for (const child of node.children) {
        dfs(child, currentLeg);
      }
    }

    // Backtrack by removing the last node in the current leg
    currentLeg.pop();
  }

  dfs(root, []);

  return result;
}

// Create the tree
const tree = new TreeNode(1);
tree.children.push(new TreeNode(2));
const node3 = new TreeNode(3);
node3.children.push(new TreeNode(4));
node3.children.push(new TreeNode(5));
node3.children.push(new TreeNode(6));
node3.children.push(new TreeNode(7));
tree.children.push(node3);
const node10 = new TreeNode(10);
tree.children.push(node10);

// Get all leg combinations
const legCombinations = getLegCombinations(tree);

// Print the result
console.log(legCombinations);

// Function to get all leaf nodes below a given node
function getLeafNodes(node) {
  const result = [];

  function dfs(currentNode) {
    if (!currentNode) return;

    if (currentNode.children.length === 0) {
      // Leaf node, add to the result
      result.push(currentNode);
    } else {
      // Non-leaf node, continue the traversal
      for (const child of currentNode.children) {
        dfs(child);
      }
    }
  }

  dfs(node);

  return result;
}
