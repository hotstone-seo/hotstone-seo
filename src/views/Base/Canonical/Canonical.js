import React, { Component } from 'react';
import {  Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Row, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Canonical extends Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    const { history } = this.props;
    history.push('/base/CanonicalForm');
  }
  render() {
    return (

      <div className="animated fadeIn">       
        <Row>
        <Col xs="6" lg="2">
        <Button block color="primary" onClick={this.handleClick} >Add New</Button>
        
            </Col>
            
          <Col xs="12" lg="12">
            <Card>
              <CardHeader>
                Canonical
              </CardHeader>
              <CardBody>
                <Table responsive bordered>
                  <thead>
                  <tr>
                    <th>Canonical Tag</th>
                    <th>Rule</th>
                     
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr>
                     <td>http://tiket.com/asad</td>
                    <td>Airport</td>
                    
                     
                    <th>Nov 15 2019</th>
                    <td>
                    <NavLink href="#">Edit</NavLink>
                    </td>
                  </tr>
                 
                  </tbody>
                </Table>
                <Pagination>
                  <PaginationItem><PaginationLink previous tag="button">Prev</PaginationLink></PaginationItem>
                  <PaginationItem active>
                    <PaginationLink tag="button">1</PaginationLink>
                  </PaginationItem>
                  
                  <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
                </Pagination>
              </CardBody>
            </Card>
         </Col>
        </Row>
      </div>
    );
  }
}
Canonical.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default Canonical;
