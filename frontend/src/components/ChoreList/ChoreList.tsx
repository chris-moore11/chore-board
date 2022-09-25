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
      description
      image
    }
    users {
	  id
      name
      email
      demerits
      choreId
      admin
    }
  }
`;

interface Chore {
	id: number
	text: string
	done: boolean
	description: string
	image: string
}

interface User {
	id: number
  	name: string
  	email: string
  	demerits: number
  	admin: boolean
  	choreId: number
}

export function ChoreList(props: IProps) {
	const queryParams = new URLSearchParams(window.location.search);
	const admin = queryParams.get('admin');
	const { data, loading, error } = useQuery(CHORES_USERS_COMBINED_QUERY, {
	  fetchPolicy: 'network-only',
	  nextFetchPolicy: 'network-only'
	});

	if (loading) {
		return <b>Loading...</b>
	}

	const findUsers = (chore: Chore) => 
		data.users
			.filter((user: User) => user.choreId == chore.id)
			.map((user: User) => user.name)
			.join(', ');


	return (
		<div className='choreList'>
			<ol>
				{data.chores.map((chore: Chore) => (
					<li key={chore.id}>
						<span className='chore'>{chore.text}:</span>
						<span className='user'>{findUsers(chore)}</span>
					</li>
				))}
		    </ol>
		    {admin && 
		     	<div className='buttons'>
		     		<RotateBackward/>
		     		<RotateForward/>
		     	</div>
	     	}
		</div>
	)
}
