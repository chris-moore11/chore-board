import React from 'react';
import logo from './logo.svg';
import './App.css';
import { useQuery, gql } from "@apollo/client";

const CREATE_CHORE_QUERY = gql`
  {
    createChore(input: { text: "chore", userId: "1" }) {
      user {
        id
      }
      text
      done
    }
  }
`;

function App() {
  return (
    <button onClick={log}>God fucking damnit bart!</button>
  );

  function log() {
    // 'must be in a react component'
    const { data, loading, error } = useQuery(CREATE_CHORE_QUERY);

    if (loading) return "Loading...";
    if (error) return error.message;

    console.log(data);
  }
}

export default App;
