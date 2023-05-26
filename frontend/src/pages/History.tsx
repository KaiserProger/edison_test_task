import React, { useEffect, useState } from 'react';
import { UserNumberCard } from '../components/UserNumberCard';
import { Col, Container, Row } from 'react-bootstrap';
import { transactionHistory } from '../network/transactions';
import { TransactionHistory } from '../types/dto/transaction_history';
import { Navigation } from '../components/Navigation';

export const History = () => {
  const [transactions, setTransactions] = useState<TransactionHistory[]>([]);
  useEffect(() => {
    const handler = async () => {
      await transactionHistory();
    };
    handler();
    setTransactions(JSON.parse(sessionStorage.getItem("transaction-history") || "[]"));
  }, []);
  return (
    <>
      <Navigation/>
      <Container>
        <Row>
          <Col className='d-flex flex-row'>
            {transactions.map((value, index) => {
              return <UserNumberCard {...value} key={index}></UserNumberCard>
            })}
          </Col>
        </Row>
      </Container>
    </>
  );
};
