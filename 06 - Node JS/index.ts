import { initializeApp } from "firebase/app";
import { getFirestore, collection, getDocs } from "firebase/firestore";

// TODO: Replace the following with your app's Firebase project configuration
// See: https://support.google.com/firebase/answer/7015592
const firebaseConfig = {
  apiKey: "AIzaSyDKTEzgxCfOESxkWsPizHux-h-xzTr1-Yw",
  authDomain: "internal-dev-client.firebaseapp.com",
  projectId: "internal-dev-client",
  storageBucket: "internal-dev-client.appspot.com",
  messagingSenderId: "5111030859",
  appId: "1:5111030859:web:ff0df8a37961f2391419ce",
  measurementId: "G-Z23QCP5P69",
};
//{
//   apiKey: "AIzaSyCaOfcjCQYDn54CYZBp40nB8lZKVB9i1GI",
//   authDomain: "beejek-67764.firebaseapp.com",
//   projectId: "beejek-67764",
//   storageBucket: "beejek-67764.appspot.com",
//   messagingSenderId: "498651717107",
//   appId: "1:498651717107:web:3577017fb81a9377a1c686",
//   measurementId: "G-CDXZYCRX69",
// };

// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Initialize Cloud Firestore and get a reference to the service
const db = getFirestore(app);

getDocs(collection(db, "jDVCc97TKqL")).then(querySnapshot => {
  querySnapshot.forEach(doc => {
    // doc.data() is never undefined for query doc snapshots
    console.log(doc.id, " => ", doc.data());
  });
});
