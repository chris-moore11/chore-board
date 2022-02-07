import React from 'react';
import { gql, useMutation } from '@apollo/client';
import { CHORES_USERS_COMBINED_QUERY } from  './ChoreList';
import { AiOutlineRight } from 'react-icons/ai'

require('./rotateForward.css');

export const FORWARD_QUERY = gql `
  mutation rotateForward {
    rotateForward
  }
`;

export function RotateForward() {
	const [onClickHandler] = useMutation(FORWARD_QUERY, {
		refetchQueries: [{query: CHORES_USERS_COMBINED_QUERY}]
	});
	const text = "->";
	return (
		<button className="forward" onClick={() => onClickHandler()}><AiOutlineRight size={28}/></button>
	)
}
