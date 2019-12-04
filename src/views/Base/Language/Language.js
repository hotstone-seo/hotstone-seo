import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Language extends Component {
    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this);
    }
    handleClick() {
        const { history } = this.props;
        history.push('/base/LanguageForm');
    }
    render() {
        return (
            <div className="animated fadeIn">
                <Col xs="12" lg="12">
                    <Card>
                        <CardHeader>
                            Language
                        </CardHeader>
                        <CardBody>
                            <div style={{ marginBottom: '.5rem' }}>
                                <Button color="primary" onClick={this.handleClick}>Add New</Button>
                            </div>
                            <Table responsive bordered>
                                <thead>
                                    <tr>
                                        <th>Name</th>
                                        <th>Language Code</th>
                                        <th>Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <td>Indonesia</td>
                                        <td>ID</td>
                                        <td><NavLink href="#">Edit</NavLink></td>
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
            </div>
        );
    }
}
Language.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};

export default Language;
