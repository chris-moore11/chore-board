import React from 'react';
import { gql, useMutation } from '@apollo/client';
import { CHORES_USERS_COMBINED_QUERY } from  './ChoreList';
import { AiOutlineLeft } from 'react-icons/ai'

require('./rotateBackward.css');

export const BACKWARD_QUERY = gql `
  mutation rotateBackward {
    rotateBackward
  }
`;

export function RotateBackward() {
	const [onClickHandler] = useMutation(BACKWARD_QUERY, {
		refetchQueries: [{query: CHORES_USERS_COMBINED_QUERY}]
	});
	const text = "<-";
	return (
		<button className="backward" onClick={() => onClickHandler()}><AiOutlineLeft size={28}/></button>
	)
}
