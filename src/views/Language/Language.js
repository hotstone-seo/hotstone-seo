import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Language extends Component {
    constructor(props) {
        super(props);
        this.state = {
            languages: [],
            record: {},
            modal: false,
            warning: false,
            formVisible: false,
            actionForm: "",
            languageFormValues: {
                id: null,
                name: null,
                code: null
            },
            URL_API: process.env.REACT_APP_API_URL + 'languages'
        }
        this.handleClick = this.handleClick.bind(this);
        this.handleEdit = this.handleEdit.bind(this);
    }

    toggle() {
        this.setState({
            modal: !this.state.modal,
        });
    }
    toggleWarning() {
        this.setState({
            warning: !this.state.warning,
        });
    }
    getLanguageList() {
        axios.get(this.state.URL_API)
            .then((res) => {
                const languages = res.data;
                this.setState({ languages });
            }).catch((error) => {

            });
    }
    componentDidMount() {
        this.getLanguageList();
    }
    handleDelete(id) {
        axios.delete(this.state.URL_API + `/${id}`)
            .then(() => {
                const { languages } = this.state;
                this.setState({ languages: languages.filter((rul) => rul.id !== id) });
            })
            .catch((error) => {

            });
        this.toggleWarning()
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
        this.setState({ formVisible: true });
    }

    saveFormRef(formRef) {
        this.formRef = formRef;
    }
    handleCancel() {
        this.setState({ formVisible: false });
    }

    handleSave() {
        const { languageFormValues, languages, actionForm, record } = this.state;
        const isUpdate = actionForm !== "Add";

        languageFormValues.id = record.id;

        if (isUpdate) {
            axios.put(this.state.URL_API, languageFormValues)
                .then(() => {
                    const index = languages.findIndex((rul) => rul.id === record.id);
                    if (index > -1) {
                        languages[index] = languageFormValues;
                        this.setState({ languages });
                    }
                })
                .then(() => {
                    this.getLanguageList();
                })
                .catch((error) => {
                    console.log(error.message)
                });
        }
        else {
            axios.post(this.state.URL_API, languageFormValues)
                .then((response) => {
                    this.setState({ languages: [...languages, languageFormValues] });
                })
                .then(() => {
                    this.getLanguageList();
                })
                .catch((error) => {

                });
            this.setState({ languageFormValues: {} });
        }
        this.setState({ formVisible: false });
    }

    handleOnChange(type, e) {
        const { target } = e || {};
        const { value } = target || {};
        const { languageFormValues } = this.state;

        this.setState({
            languageFormValues: {
                ...languageFormValues,
                [type]: value
            }
        });
    }

    /*handleClick() {
        const { history } = this.props;
        history.push('/languageForm');
    }
    handleEdit() {
        const { history } = this.props;
        history.push('/languageEditForm');
    }*/
    render() {
        const { languages } = this.state;
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
                                        <td><NavLink href="#" onClick={this.handleEdit}>Edit</NavLink></td>
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
