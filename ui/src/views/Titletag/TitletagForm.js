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

class TitletagForm extends Component {
    constructor(props) {
        super(props);
        this.toggle = this.toggle.bind(this);
        this.toggleFade = this.toggleFade.bind(this);
        this.state = {
            collapse: true,
            fadeIn: true,
            timeout: 300
        };
        this.handlePreview = this.handlePreview.bind(this);
    }

    toggle() {
        this.setState({ collapse: !this.state.collapse });
    }

    toggleFade() {
        this.setState((prevState) => { return { fadeIn: !prevState } });
    }

    handlePreview() {
        const { history } = this.props;
        history.push('/titletagPreview');
    }

    render() {
        const {
            visible, onCancel, onSave, titletag, action, onChange
        } = this.props;
        return (
            <Modal isOpen={visible}>
                <ModalHeader>{action} Title-Tag</ModalHeader>
                <ModalBody>
                    <Form className="form-horizontal">
                        <FormGroup row>
                            <Col md="3">
                                <Label htmlFor="text-input">Language</Label>
                            </Col>
                            <Col xs="12" md="9">
                                {titletag !== undefined ? (<Input type="hidden" id="id" name="id" defaultValue={titletag !== undefined ? titletag.id : ""} onChange={onChange.bind(this, 'id')} />) : ""}
                                <Input type="select" name="language" id="language">
                                    <option>ID</option>
                                </Input>
                            </Col>
                        </FormGroup>
                        <FormGroup row>
                            <Col md="3">
                                <Label htmlFor="text-input">Title</Label>
                            </Col>
                            <Col xs="12" md="9">
                                <Input
                                    type="text"
                                    id="title"
                                    name="title"
                                    placeholder="title"
                                    defaultValue={titletag !== undefined ? titletag.title : ""}
                                    onChange={onChange.bind(this, 'title')}
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
TitletagForm.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};
export default TitletagForm;