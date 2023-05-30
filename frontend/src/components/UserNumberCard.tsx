import React from 'react';
import { TransactionHistory } from '../types/dto/transaction_history';


export const UserNumberRow = (props: TransactionHistory) => {
  return (
    <tr>
      <td>{props.correct_number}</td>
      {props.extrasense_guesses.map((value) => (
        <td>{value.number}</td>
      ))}
    </tr>
  )
}