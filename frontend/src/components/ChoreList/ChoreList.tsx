import React from 'react';
import { useState } from 'react';
import { gql, useQuery } from '@apollo/client';
import { RotateForward } from './RotateForward';
import { RotateBackward } from './RotateBackward';

require('./choreList.css');

interface IProps {}

export const CHORES_USERS_COMBINED_QUERY = gql`
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
	const queryParams = new URLSearchParams(window.location.search);
	const admin = queryParams.get('admin');
	const { data, loading, error } = useQuery(CHORES_USERS_COMBINED_QUERY, {
	  fetchPolicy: "network-only",
	  nextFetchPolicy: "network-only"
	});

	if (loading) {
		return <b>Loading...</b>
	}

	return (
		<div className="choreList">
			<ol>
		        {data.chores.map((chore: Chore) => (
		          <li key={chore.id}>
		          	<span className="chore">{data.chores[chore.id-1].text}:</span>
		          	<span className="user">{data.users.find((user: User) => user.choreId === chore.id).name}</span>
	          	  </li>
		        ))}
		    </ol>
		    {admin && 
		     	<div className="buttons">
		     		<RotateBackward/>
		     		<RotateForward/>
		     	</div>
	     	}
		</div>
	)
}

