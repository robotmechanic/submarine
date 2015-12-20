﻿// This file was generated by typhen-api

using System;

namespace TyphenApi
{
    public interface IWebApi<ErrorT> where ErrorT : TypeBase, new()
    {
        IWebApiRequestSender RequestSender { get; }
        ISerializer RequestSerializer { get; }
        IDeserializer ResponseDeserializer { get; }

        void OnBeforeRequestSend(IWebApiRequest request);
        void OnRequestError(IWebApiRequest request, WebApiError<ErrorT> error);
        void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response);
    }

    public abstract class WebApiBase<ErrorT> : IWebApi<ErrorT> where ErrorT : TypeBase, new()
    {
        public Uri BaseUri { get; private set; }

        public IWebApiRequestSender RequestSender { get; protected set; }
        public ISerializer RequestSerializer { get; protected set; }
        public IDeserializer ResponseDeserializer { get; protected set; }

        public abstract void OnBeforeRequestSend(IWebApiRequest request);
        public abstract void OnRequestError(IWebApiRequest request, WebApiError<ErrorT> error);
        public abstract void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response);

        protected WebApiBase(string baseUri)
        {
            BaseUri = new Uri(baseUri);
        }

        public WebApiError<ErrorT> Error(Exception error)
        {
            return error as WebApiError<ErrorT>;
        }
    }
}
