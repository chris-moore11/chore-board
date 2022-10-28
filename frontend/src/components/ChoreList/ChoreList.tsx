import React from 'react';
import { useState } from 'react';
import ReactTooltip from 'react-tooltip';
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
	const usersView = queryParams.get('usersView');
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
			{usersView ?
				<ol>
					{data.users.map((user: User) => (
						<li key={user.id}>
							<span className='user'>{user.name}:</span>
							<span className='chore' data-tip={data.chores.find((chore: Chore) => chore.id == user.choreId).description} data-for='tooltip'>
								{data.chores.find((chore: Chore) => chore.id == user.choreId).text}
							</span>
							<ReactTooltip id='tooltip'/>
						</li>
					))}
		    	</ol> :
				<ol>
					{data.chores.map((chore: Chore) => (
						<li key={chore.id}>
							<span className='chore' data-tip={chore.description} data-for='tooltip'>{chore.text}:</span>
							<ReactTooltip id='tooltip'/>
							<span className='user'>{findUsers(chore)}</span>
						</li>
					))}
				</ol>
			}
		    {admin && 
		     	<div className='buttons'>
		     		<RotateBackward/>
		     		<RotateForward/>
		     	</div>
	     	}
		</div>
	)
}
