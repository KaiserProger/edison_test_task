import React, { useEffect, useState } from 'react';
import { UserNumberRow } from '../components/UserNumberCard';
import { Col, Container, Row, Table } from 'react-bootstrap';
import { transactionHistory } from '../network/transactions';
import { TransactionHistory } from '../types/dto/transaction_history';
import { Navigation } from '../components/Navigation';

export const History = () => {
  const [transactions, setTransactions] = useState<TransactionHistory[]>([]);
  useEffect(() => {
    const handler = async () => {
      await transactionHistory();
      setTransactions(JSON.parse(sessionStorage.getItem("transaction-history") || "[]"));
    };
    handler();
  }, []);
  return (
    <>
      <Navigation/>
      <Container>
        <Row>
          <Col className='d-flex flex-row'>
            <Table>
              <thead>
                <tr>
                  <th>Number</th>
                  {transactions.length === 0 && (
                    <h1>There's no numbers yet...</h1>
                  )}
                  {transactions.length > 0 && transactions[0].extrasense_guesses.map((value) => (
                    <th style={{
                        "wordWrap": "break-word", 
                        "minWidth": "160px",
                        "maxWidth": "160px"
                      }}>{value.guessed_by}</th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {transactions.map((value, index) => (
                  <UserNumberRow {...value} key={index} />
                ))}
              </tbody>
            </Table>
          </Col>
        </Row>
      </Container>
    </>
  );
};
