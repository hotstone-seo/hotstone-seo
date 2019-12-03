import React, { Component } from 'react';
import {  Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Row, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Titletag extends Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    const { history } = this.props;
    history.push('/base/TitletagForm');
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
                Title-Tag
              </CardHeader>
              <CardBody>
                <Table responsive bordered>
                  <thead>
                  <tr>
                    <th>Rule</th>
                    <th>Language</th>
                    <th>Title</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr>
                    <td>Airport Detail</td>
                    <td>ID</td>
                    <td>.. is located at ..</td>
                    <th>Nov 16 2019</th>
                    <td>
                    <NavLink href="#">Edit</NavLink>
                    </td>
                  </tr>
                  <tr>
                     <td>Airport Detail</td>
                    <td>ENG</td>
                    <td>.. is located at ..</td>
                    <th>Nov 16 2019</th>
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
                  <PaginationItem className="page-item"><PaginationLink tag="button">2</PaginationLink></PaginationItem>
                  <PaginationItem><PaginationLink tag="button">3</PaginationLink></PaginationItem>
                  <PaginationItem><PaginationLink tag="button">4</PaginationLink></PaginationItem>
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
Titletag.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default Titletag;
