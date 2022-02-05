import React from 'react';
import './App.css';
import { useQuery, gql } from "@apollo/client";

const CREATE_CHORE_QUERY = gql`
  {
    chores {
		text
		done
		user {
  			name
		}
	}
  }
`;

interface User {
	name: string,
}

interface Chore {
	text: string,
	done: boolean,
	user: User,
}

function App() {
  const { data, loading, error } = useQuery(CREATE_CHORE_QUERY);

  if (loading) return <b>"Loading..."</b>;
  if (error) return <pre>{error.message}</pre>

  //return (
  // <button onClick={log}>God fucking damnit bart!</button>
  //);
//
  //function log() {
  //  console.log('bill');
  //}

  return (
    <div>
      <h1>SpaceX Launches</h1>
      <ul>
        {data.chores.map((chore: Chore) => (
          <li key={chore.text}>{chore.text}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
