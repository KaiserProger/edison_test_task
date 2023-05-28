import React, { useEffect, useState } from 'react'
import { Col, Container, Row } from 'react-bootstrap'
import { TrustLevel } from '../types/dto/trust_level';
import { fetchTrustLevels } from '../network/trust';
import { Navigation } from '../components/Navigation';

export const ExtrasenseTrustPage = () => {
  const [trustLevels, setTrustLevels] = useState<TrustLevel[]>([]);
  useEffect(() => {
    const handler = async () => {
      await fetchTrustLevels();
      setTrustLevels(JSON.parse(sessionStorage.getItem("trust-levels") || "[]"));
    };
    handler();
  }, [setTrustLevels]);
  return (
    <>
      <Navigation/>
      <Container>
        <Row>
          <Col>
            {trustLevels.map((value) => {
              return (<p>{value.extrasenseId}:{value.trust}</p>);
            })}
          </Col>
        </Row>
      </Container>
    </>
  );
};
