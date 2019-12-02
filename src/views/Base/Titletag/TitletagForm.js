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
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
  Input,
  InputGroup,
  InputGroupButtonDropdown,
  Label,
  Row,
} from 'reactstrap';
import Scripttag from './Scripttag';

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
    this.setState((prevState) => { return { fadeIn: !prevState }});
  }

  handlePreview() {
    const { history } = this.props;
    history.push('/base/TitletagPreview');
  }

  render() {
    return (
      <div className="animated fadeIn">
     
        <Row>
          <Col xs="12" md="12">
            <Card>
              <CardHeader>
                <strong>Add New Title-Tag</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Rule</Label>
                    </Col>
                    <Col xs="12" md="9">
                        <InputGroup>
                        <InputGroupButtonDropdown addonType="prepend"
                                                  isOpen={this.state.first}
                                                  toggle={() => { this.setState({ first: !this.state.first }); }}>
                          <DropdownToggle caret color="primary">
                            -Choose-
                          </DropdownToggle>
                          <DropdownMenu className={this.state.first ? 'show' : ''}>
                            <DropdownItem>Airport Detail</DropdownItem>                           
                          </DropdownMenu>
                        </InputGroupButtonDropdown>                     
                      </InputGroup>
                    </Col>
                  </FormGroup>

                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Language</Label>
                    </Col>
                    <Col xs="12" md="9">
                      
                       
                    </Col>
                  </FormGroup>

                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Title</Label>
                    </Col>
                    <Col xs="12" md="9">
                         
                    </Col>
                  </FormGroup>

                </Form>
              </CardBody>
              <CardFooter>
                <Button type="submit" size="sm" color="primary"><i className="fa fa-dot-circle-o"></i> Submit</Button>
                <Button type="button" size="sm" color="secondary" onClick={this.handlePreview}><i className="fa fa-ban"></i> Preview</Button>
              </CardFooter>
            </Card>
            
          </Col>
        </Row>         
      </div>
    );
  }
}
TitletagForm.propTypes = {
    match: PropTypes.shape({
      path: PropTypes.string,
    }).isRequired,
};
export default TitletagForm;