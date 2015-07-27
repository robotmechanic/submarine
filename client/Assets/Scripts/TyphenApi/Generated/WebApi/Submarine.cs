// This file was generated by typhen-api

using System;

namespace TyphenApi.WebApi
{
    public class Submarine : TyphenApi.WebApiBase<TyphenApi.Type.Submarine.Error>
    {
        public Submarine(string baseUri) : base(baseUri)
        {
            Controller = new TyphenApi.Controller.WebApi.Submarine();
        }

        public Submarine(string baseUri, IWebApiController<TyphenApi.Type.Submarine.Error> controller) : base(baseUri)
        {
            Controller = controller;
        }

        class PingRequestParameters : TyphenApi.TypeBase<PingRequestParameters>
        {
            protected static readonly SerializationInfo<PingRequestParameters, string> message = new SerializationInfo<PingRequestParameters, string>("message", false, (x) => x.Message, (x, v) => x.Message = v);
            public string Message { get; set; }
        }

        public TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.PingObject, TyphenApi.Type.Submarine.Error> Ping(string message)
        {
            var requestBody = new PingRequestParameters();
            requestBody.Message = message;

            var request = new TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.PingObject, TyphenApi.Type.Submarine.Error>(Controller);
            request.Uri = new Uri(BaseUri, "ping");
            request.Method = HttpMethod.Post;
            request.Body = requestBody;
            request.NoAuthenticationRequired = true;
            return request;
        }

        class SignUpRequestParameters : TyphenApi.TypeBase<SignUpRequestParameters>
        {
            protected static readonly SerializationInfo<SignUpRequestParameters, string> name = new SerializationInfo<SignUpRequestParameters, string>("name", false, (x) => x.Name, (x, v) => x.Name = v);
            public string Name { get; set; }
            protected static readonly SerializationInfo<SignUpRequestParameters, string> password = new SerializationInfo<SignUpRequestParameters, string>("password", false, (x) => x.Password, (x, v) => x.Password = v);
            public string Password { get; set; }
        }

        public TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.SignUpObject, TyphenApi.Type.Submarine.Error> SignUp(string name, string password)
        {
            var requestBody = new SignUpRequestParameters();
            requestBody.Name = name;
            requestBody.Password = password;

            var request = new TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.SignUpObject, TyphenApi.Type.Submarine.Error>(Controller);
            request.Uri = new Uri(BaseUri, "sign_up");
            request.Method = HttpMethod.Post;
            request.Body = requestBody;
            request.NoAuthenticationRequired = true;
            return request;
        }

        class LoginRequestParameters : TyphenApi.TypeBase<LoginRequestParameters>
        {
            protected static readonly SerializationInfo<LoginRequestParameters, string> name = new SerializationInfo<LoginRequestParameters, string>("name", false, (x) => x.Name, (x, v) => x.Name = v);
            public string Name { get; set; }
            protected static readonly SerializationInfo<LoginRequestParameters, string> password = new SerializationInfo<LoginRequestParameters, string>("password", false, (x) => x.Password, (x, v) => x.Password = v);
            public string Password { get; set; }
        }

        public TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.LoginObject, TyphenApi.Type.Submarine.Error> Login(string name, string password)
        {
            var requestBody = new LoginRequestParameters();
            requestBody.Name = name;
            requestBody.Password = password;

            var request = new TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.LoginObject, TyphenApi.Type.Submarine.Error>(Controller);
            request.Uri = new Uri(BaseUri, "login");
            request.Method = HttpMethod.Post;
            request.Body = requestBody;
            request.NoAuthenticationRequired = true;
            return request;
        }

        class FindUserRequestParameters : TyphenApi.TypeBase<FindUserRequestParameters>
        {
            protected static readonly SerializationInfo<FindUserRequestParameters, string> name = new SerializationInfo<FindUserRequestParameters, string>("name", false, (x) => x.Name, (x, v) => x.Name = v);
            public string Name { get; set; }
        }

        public TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.FindUserObject, TyphenApi.Type.Submarine.Error> FindUser(string name)
        {
            var requestBody = new FindUserRequestParameters();
            requestBody.Name = name;

            var request = new TyphenApi.WebApiRequest<TyphenApi.Type.Submarine.FindUserObject, TyphenApi.Type.Submarine.Error>(Controller);
            request.Uri = new Uri(BaseUri, "find_user");
            request.Method = HttpMethod.Post;
            request.Body = requestBody;
            request.NoAuthenticationRequired = false;
            return request;
        }
    }
}
