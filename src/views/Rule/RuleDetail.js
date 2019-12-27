import React, { Component } from 'react';
import {
    Button,
    Card,
    CardBody,
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

import CanonicalForm from '../Canonical/CanonicalForm';
import MetaTagForm from '../Metatag/MetatagForm';
import ScriptTagForm from '../Scripttag/ScripttagForm';
import TitleTagForm from '../Titletag/TitletagForm';

import axios from 'axios';

export const parseQuery = (subject) => {
    const results = {};
    const parser = /[^&?]+/g;
    let match = parser.exec(subject);
    while (match !== null) {
        const parts = match[0].split('=');
        results[parts[0]] = parts[1];
        match = parser.exec(subject);
    }
    return results;
};


class RuleDetail extends Component {
    constructor(props) {
        super(props);

        this.toggle = this.toggle.bind(this);
        this.toggleFade = this.toggleFade.bind(this);
        this.state = {
            collapse: true,
            fadeIn: true,
            timeout: 300,
            URL_API: process.env.REACT_APP_API_URL + 'rules',
            canonicalFormValues: {
                id: null,
                name: null,
                rule_id: null,
            },
            metaTagFormValues: {
                id: null,
                name: null,
                rule_id: null,
            },
            scriptTagFormValues: {
                id: null,
                name: null,
                rule_id: null,
            },
            titleTagFormValues: {
                id: null,
                name: null,
                rule_id: null,
            },
            canonicalFormVisible: false,
            actionCanonicalForm: "",
            metaTagFormVisible: false,
            scriptTagFormVisible: false,
            titleTagFormVisible: false,
            ruleIdParam: 0,
        };
        this.handleEditCanonical = this.handleEditCanonical.bind(this);

        this.handleAddNewCanonical = this.handleAddNewCanonical.bind(this);
        this.handleAddNewMeta = this.handleAddNewMeta.bind(this);
        this.handleAddNewScript = this.handleAddNewScript.bind(this);
        this.handleAddNewTitle = this.handleAddNewTitle.bind(this);
        this.handleCancelAddCanonical = this.handleCancelAddCanonical.bind(this);
        this.handleCancelAddMetaTag = this.handleCancelAddMetaTag.bind(this);
        this.handleCancelAddScriptTag = this.handleCancelAddScriptTag.bind(this);
        this.handleCancelAddTitleTag = this.handleCancelAddTitleTag.bind(this);
    }
    componentDidMount() {
        //const query = parseQuery((window.location || {}).search || '');

        const { ruleId } = this.state;


    }

    toggle() {
        this.setState({ collapse: !this.state.collapse });
    }

    toggleFade() {
        this.setState((prevState) => { return { fadeIn: !prevState } });
    }

    handleEditCanonical() {
        const { history } = this.props;
        history.push('/canonicalEditForm');
    }
    handleAddNewCanonical() {
        const { history } = this.props;
        history.push('/canonicalForm');
    }
    handleAddNewMeta() {
        const { history } = this.props;
        history.push('/metatagForm');
    }
    handleAddNewScript() {
        const { history } = this.props;
        history.push('/scripttagForm');
    }
    handleAddNewTitle() {
        const { history } = this.props;
        history.push('/titletagForm');
    }
    showForm(record) {
        if (record !== undefined) {
            this.setState({ record: record });
            this.setState({ actionForm: "Edit" });
        }
        else {
            this.setState({ record: {} });
            this.setState({ actionForm: "Add" });
        }
        this.setState({ canonicalFormVisible: true });
    }
    showFormMetaTag(record) {
        if (record !== undefined) {
            this.setState({ record: record });
            this.setState({ actionForm: "Edit" });
        }
        else {
            this.setState({ record: {} });
            this.setState({ actionForm: "Add" });
        }
        this.setState({ metaTagFormVisible: true });
    }

    showFormScriptTag(record) {
        if (record !== undefined) {
            this.setState({ record: record });
            this.setState({ actionForm: "Edit" });
        }
        else {
            this.setState({ record: {} });
            this.setState({ actionForm: "Add" });
        }
        this.setState({ scriptTagFormVisible: true });
    }

    showFormTitleTag(record) {
        if (record !== undefined) {
            this.setState({ record: record });
            this.setState({ actionForm: "Edit" });
        }
        else {
            this.setState({ record: {} });
            this.setState({ actionForm: "Add" });
        }
        this.setState({ titleTagFormVisible: true });
    }

    handleOnChange(type, e) {
        const { target } = e || {};
        const { value } = target || {};
        const { canonicalFormValues } = this.state;

        this.setState({
            canonicalFormValues: {
                ...canonicalFormValues,
                [type]: value
            }
        });
    }
    handleCancelAddCanonical() {
        this.setState({ canonicalFormVisible: false });
    }
    handleCancelAddMetaTag() {
        this.setState({ metaTagFormVisible: false });
    }
    handleCancelAddScriptTag() {
        this.setState({ scriptTagFormVisible: false });
    }
    handleCancelAddTitleTag() {
        this.setState({ titleTagFormVisible: false });
    }
    render() {
        const query = parseQuery((window.location || {}).search || '');
        const { ruleId } = query || {};

        axios.get(this.state.URL_API + `/${ruleId}`)
            .then((res) => {
                const rulesdata = res.data;
                console.log(rulesdata, "rules");
                //console.log(ruleId,"rules id");
                this.setState({ rules, rulesdata });
            }).catch((error) => {
                //this.toggleWarningAPI(error.message)
            });

        //const { data } = this.props.location;
        const { rules } = this.state;
        
        return (
            <div className="animated fadeIn">
                 
                        <Row>
                            <Col xs="12" md="9" lg="6">
                                <Card>
                                    <CardHeader>
                                        <strong>Detail Rule ID {rules.id}</strong>
                                    </CardHeader>
                                    <CardBody>
                                        <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                                            <FormGroup row>
                                                <Col md="3">
                                                    <Label htmlFor="text-input">Name</Label>
                                                </Col>
                                                <Col xs="12" md="9">
                                                    {rules.name}
                                                </Col>
                                            </FormGroup>
                                            <FormGroup row>
                                                <Col md="3">
                                                    <Label htmlFor="text-input">URL Pattern</Label>
                                                </Col>
                                                <Col xs="12" md="9">
                                                    {rules.url_pattern}
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
                                </Card>
                            </Col>
                        </Row>
                   


                <Row>
                    <Col>
                        <Card>
                            <CardHeader>
                                <i className="fa fa-align-justify"></i>
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={() => this.showForm()} style={{ marginRight: "0.4em" }}>Add New Canonical</Button>
                                    <Button color="primary" onClick={() => this.showFormMetaTag()} style={{ marginRight: "0.4em" }}>Add New Meta-Tag</Button>
                                    <Button color="primary" onClick={() => this.showFormScriptTag()} style={{ marginRight: "0.4em" }}>Add New Script Tag</Button>
                                    <Button color="primary" onClick={() => this.showFormTitleTag()} style={{ marginRight: "0.4em" }}>Add New Title-Tag</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>
                                            <th>Type</th>
                                            <th>Attribute</th>
                                            <th>Value</th>
                                            <td>Language</td>
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td>Canonical</td>
                                            <td>xxx</td>
                                            <td>http://tiket.com/asad</td>
                                            <td>ID</td>
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
                        <CanonicalForm
                            visible={this.state.canonicalFormVisible}
                            onCancel={this.handleCancelAddCanonical}
                            onSave={this.handleSave}
                            canonical={this.state.record}
                            action={this.state.actionForm}
                            onChange={this.handleOnChange.bind(this)}
                        />
                        <MetaTagForm
                            visible={this.state.metaTagFormVisible}
                            onCancel={this.handleCancelAddMetaTag}
                            onSave={this.handleSave}
                            metatag={this.state.record}
                            action={this.state.actionForm}
                            onChange={this.handleOnChange.bind(this)}
                        />
                        <ScriptTagForm
                            visible={this.state.scriptTagFormVisible}
                            onCancel={this.handleCancelAddScriptTag}
                            onSave={this.handleSave}
                            scripttag={this.state.record}
                            action={this.state.actionForm}
                            onChange={this.handleOnChange.bind(this)}
                        />
                        <TitleTagForm
                            visible={this.state.titleTagFormVisible}
                            onCancel={this.handleCancelAddTitleTag}
                            onSave={this.handleSave}
                            titletag={this.state.record}
                            action={this.state.actionForm}
                            onChange={this.handleOnChange.bind(this)}
                        />
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