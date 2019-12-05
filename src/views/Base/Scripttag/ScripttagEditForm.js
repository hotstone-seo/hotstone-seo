import React, { Component } from 'react';
import PropTypes from 'prop-types';

import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    Col,
    Form,
    FormGroup,
    Input,
    Label,
    Row,
} from 'reactstrap';

class ScripttagEditForm extends Component {
    constructor(props) {
        super(props);
        this.handleCancel = this.handleCancel.bind(this);
    }

    handleCancel() {
        const { history } = this.props;
        history.push('/base/scripttag');
    }

    render() {
        return (

            <div className="animated fadeIn">
                <Row>
                    <Col xs="12" md="9" lg="6">
                        <Card>
                            <CardHeader>
                                <strong>Edit Script-Tag</strong>
                            </CardHeader>
                            <CardBody>
                                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Rule</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            <Input type="select" name="rule" id="rule">
                                                <option selected>Airport Detail</option>
                                            </Input>
                                        </Col>
                                    </FormGroup>

                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Type</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            <Input type="select" name="type" id="type">
                                                <option selected>Javascript</option>
                                            </Input>
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Source</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            <Input type="text" id="content" name="content" placeholder="Source" value="https://www.google-anayytics.com/analytics.js" />
                                        </Col>
                                    </FormGroup>
                                </Form>
                            </CardBody>
                            <CardFooter>
                                <Button type="submit" size="md" color="primary" style={{ marginRight: "0.4em" }}><i className="fa fa-dot-circle-o"></i>Save Change</Button>
                                <Button type="button" size="md" color="secondary" onClick={this.handleCancel}>Cancel</Button>
                            </CardFooter>
                        </Card>
                    </Col>
                </Row>
            </div>
        );
    }
}
ScripttagEditForm.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};
export default ScripttagEditForm;