import React, { Component } from 'react';
import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    Col,
    Form,
    FormGroup,
    Label,
    NavLink,
    Pagination,
    PaginationItem,
    PaginationLink,
    Row,
    Table,
} from 'reactstrap';

import PropTypes from 'prop-types';


class RuleDetail extends Component {
    constructor(props) {
        super(props);

        this.toggle = this.toggle.bind(this);
        this.toggleFade = this.toggleFade.bind(this);
        this.state = {
            collapse: true,
            fadeIn: true,
            timeout: 300
        };
        this.handleBack = this.handleBack.bind(this);
        this.handleEditCanonical = this.handleEditCanonical.bind(this);
    }

    toggle() {
        this.setState({ collapse: !this.state.collapse });
    }

    toggleFade() {
        this.setState((prevState) => { return { fadeIn: !prevState } });
    }
    handleBack() {
        const { history } = this.props;
        history.push('/rule');
    }
    handleEditCanonical() {
        const { history } = this.props;
        history.push('/canonicalEditForm');
    }
    render() {
        const ruleID = this.props.match.params.id;

        return (
            <div className="animated fadeIn">
                <Row>
                    <Col xs="12" md="9" lg="6">
                        <Card>
                            <CardHeader>
                                <strong>Detail Rule ID {ruleID}</strong>
                            </CardHeader>
                            <CardBody>
                                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Name</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            Airport Detail
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">URL Pattern</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            http://xxx
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="textarea-input">Exclusion</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            -
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Data Source</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            Airport
                                        </Col>
                                    </FormGroup>
                                </Form>
                            </CardBody>
                            <CardFooter>

                                <Button type="button" size="md" color="secondary" onClick={this.handleBack}> Back</Button>
                            </CardFooter>
                        </Card>
                    </Col>
                </Row>

                <Row>
                    <Col>
                        <Card>
                            <CardHeader>
                                <i className="fa fa-align-justify"></i> Canonical
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={this.handleClick}>Add New</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>
                                            <th>Canonical Tag</th>

                                            <th>Updated Date</th>
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>http://hotstone-seo/asad</td>

                                            <td>Nov 15 2019</td>
                                            <td>
                                                <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                                <nav>
                                    <Pagination>
                                        <PaginationItem><PaginationLink previous tag="button">Prev</PaginationLink></PaginationItem>
                                        <PaginationItem active>
                                            <PaginationLink tag="button">1</PaginationLink>
                                        </PaginationItem>

                                        <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
                                    </Pagination>
                                </nav>
                            </CardBody>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col xs="12" lg="12">
                        <Card>
                            <CardHeader>
                                Meta-Tag
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={this.handleClick}>Add New</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>
                                            <th>Name</th>

                                            <th>Content</th>
                                            <th>Updated Date</th>
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>Pompeius René</td>

                                            <td>xx</td>
                                            <td>Nov 16 2019</td>
                                            <td>
                                                <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
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
                                    <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
                                </Pagination>
                            </CardBody>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col xs="12" lg="12">
                        <Card>
                            <CardHeader>
                                Title-Tag
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={this.handleClick}>Add New</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>

                                            <th>Language</th>
                                            <th>Title</th>
                                            <th>Updated Date</th>
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>

                                            <td>ID</td>
                                            <td>.. is located at ..</td>
                                            <td>Nov 16 2019</td>
                                            <td>
                                                <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
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

                <Row>
                    <Col xs="12" lg="12">
                        <Card>
                            <CardHeader>
                                Script-Tag
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={this.handleClick}>Add New</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>

                                            <th>Type</th>
                                            <th>Source</th>
                                            <th>Updated Date</th>
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>

                                            <td>Javascript</td>
                                            <td>https://www.google-anayytics.com/analytics.js</td>
                                            <td>Nov 16 2019</td>
                                            <td>
                                                <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
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
RuleDetail.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};
export default RuleDetail;