import React, { Component } from 'react';
import PropTypes from 'prop-types';

import {
  Button,
  Col,
  Form,
  FormGroup,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Input,
  Label,
} from 'reactstrap';

class CanonicalForm extends Component {
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
    this.handleCancel = this.handleCancel.bind(this);
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState((prevState) => { return { fadeIn: !prevState } });
  }

  handlePreview() {
    const { history } = this.props;
    history.push('/base/MetatagPreview');
  }

  handleCancel() {
    const { history } = this.props;
    history.push('/canonical');
  }
  render() {
    const {
      visible, onCancel, onSave, canonical, action, onChange
    } = this.props;
    return (
      <Modal isOpen={visible}>
        <ModalHeader>{action} Canonical</ModalHeader>
        <ModalBody>
          <Form className="form-horizontal">
            <FormGroup row>
              {canonical !== undefined? (<Input type="hidden" id="id" name="id" defaultValue={canonical.id} onChange={onChange.bind(this, 'id')}/>):""}
              <Col md="3">
                <Label htmlFor="text-input">Name</Label>
              </Col>
              <Col xs="12" md="9">
                <Input
                  type="text"
                  id="name"
                  name="name"
                  placeholder="Name"
                  defaultValue={canonical !== undefined?canonical.name:""}
                  onChange={onChange.bind(this, 'name')}
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
CanonicalForm.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};
export default CanonicalForm;