import React from 'react';
import { gql, useQuery } from '@apollo/client';


interface IProps {}

const CHORES_USERS_COMBINED_QUERY = gql`
  {
    chores {
	  id
      text
      done
      image
      tutorial
    }
    users {
	  id
      name
      email
      image
      choreId
      admin
    }
  }
`;

interface Chore {
	id: number,
	text: string,
	done: boolean,
	image: string,
	tutorial: string,
}

interface User {
	id: number
  	name: string
  	email: string
  	image: string
  	choreId: number
  	admin: boolean
}

export function ChoreList(props: IProps) {
	const { data, loading, error } = useQuery(CHORES_USERS_COMBINED_QUERY);

	if (loading) {
		return <b>Loading...</b>
	}

	return (
		<div>
			<ul>
		        {data.users.map((user: User) => (
		          <li key={user.id}>{user.name}: {data.chores[user.choreId-1].text}</li>
		        ))}
		     </ul>
		</div>
	)
}

