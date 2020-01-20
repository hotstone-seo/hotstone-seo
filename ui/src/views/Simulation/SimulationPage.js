import React from "react";
import { useForm } from "react-hook-form";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";

function SimulationPage() {
  const [{ data, error }, execute] = useHotstoneAPI(
    {
      url: `provider/matchRule`,
      method: "POST"
    },
    { manual: true }
  );

  const { register, handleSubmit, errors } = useForm();
  const onSubmit = ({ url }) => {
    execute({ data: { path: url } });
  };

  // inspectAxiosError(error);
  // console.log("=== DATA ===");
  // console.log(data);

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <div className="card">
            <div className="card-header">Simulation: Rule Matching</div>
            <div className="card-body">
              <form
                onSubmit={handleSubmit(onSubmit)}
                className="form-horizontal needs-validation"
              >
                <div className="form-row">
                  <div className="col-md-9">
                    <input
                      name="url"
                      placeholder="URL"
                      className={"form-control " + (errors.url && "is-invalid")}
                      ref={register({
                        required: "Required"
                      })}
                    />
                    {errors.url && (
                      <div className="invalid-feedback">
                        {errors.url.message}
                      </div>
                    )}
                  </div>
                  <div className="col-auto">
                    <button type="submit" className="btn btn-primary">
                      Submit
                    </button>
                  </div>
                </div>
              </form>
            </div>

            {renderIfSuccess(data)}
            {renderIfError(error)}
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col"></div>
        <div className="col"></div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col"></div>
      </div>
    </div>
  );
}

function renderIfSuccess(data) {
  // TODO: display path params
  if (data)
    return (
      <div className="card-footer">
        <div className="alert alert-success" role="alert">
          Matched
        </div>
      </div>
    );
}

function renderIfError(error) {
  if (error)
    return (
      <div className="card-footer">
        <div className="alert alert-danger" role="alert">
          {error.response.data.message}
        </div>
      </div>
    );
}

export default SimulationPage;
