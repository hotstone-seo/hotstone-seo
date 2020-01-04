import React, { Component } from 'react';
import PropTypes from 'prop-types';

import {
    Button,
    Col,
    Form,
    FormGroup,
    Input,
    Label,
    Modal,
    ModalBody,
    ModalFooter,
    ModalHeader,
} from 'reactstrap';


class ScripttagForm extends Component {
    constructor(props) {
        super(props);
        this.handlePreview = this.handlePreview.bind(this);
    }

    handlePreview() {
        const { history } = this.props;
        history.push('/base/ScripttagPreview');
    }

    render() {
        const {
            visible, onCancel, onSave, scripttag, action, onChange
        } = this.props;
        return (
            <Modal isOpen={visible}>
                <ModalHeader>{action} Script-Tag</ModalHeader>
                <ModalBody>
                    <Form className="form-horizontal">
                        <FormGroup row>
                            <Col md="3">
                                <Label htmlFor="text-input">Type</Label>
                            </Col>
                            <Col xs="12" md="9">
                                {scripttag !== undefined ? (<Input type="hidden" id="id" name="id" defaultValue={scripttag !== undefined ? scripttag.id : ""} onChange={onChange.bind(this, 'id')} />) : ""}
                                <Input type="select" name="type" id="type">
                                    <option>Javascript</option>
                                </Input>
                            </Col>
                        </FormGroup>
                        <FormGroup row>
                            <Col md="3">
                                <Label htmlFor="text-input">Source</Label>
                            </Col>
                            <Col xs="12" md="9">
                                <Input
                                    type="text"
                                    id="source"
                                    name="source"
                                    placeholder="Source"
                                    defaultValue={scripttag !== undefined ? scripttag.source : ""}
                                    onChange={onChange.bind(this, 'source')}
                                />
                            </Col>
                        </FormGroup>
                    </Form>
                </ModalBody>
                <ModalFooter>
                    <Button color="warning" onClick={onSave}>Save</Button>{' '}
                    <Button color="secondary" onClick={onCancel}>Cancel</Button>
                </ModalFooter>
            </Modal>
        );
    }
}
ScripttagForm.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};
export default ScripttagForm;