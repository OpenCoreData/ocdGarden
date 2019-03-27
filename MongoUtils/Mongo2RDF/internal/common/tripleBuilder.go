package common

import (
	"fmt"
	"log"
	"strings"

	"github.com/knakk/rdf"
)

// SPOIRI return a new triple composed or two IRI's
func SPOIRI(subj, pred, obj string) rdf.Triple {
	newsub, err := rdf.NewIRI(subj)
	newpred0, err := rdf.NewIRI(pred)
	newobj0, err := rdf.NewIRI(obj)
	newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

	if err != nil {
		log.Printf("this is error %v \n", err)
	}

	return newtriple0
}

// SPOLiteral return a new triple with literial object
func SPOLiteral(subj, pred, obj string) rdf.Triple {
	newsub, err := rdf.NewIRI(subj)
	newpred0, err := rdf.NewIRI(pred)
	newobj0, err := rdf.NewLiteral(obj)
	newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}

	if err != nil {
		log.Printf("this is error %v \n", err)
	}

	return newtriple0
}

// IITriple for IRI IRI sets
func IITriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	if o == "" {
		return nil
	}

	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewIRI(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	fmt.Fprintf(b, "%s", qs)
	return err
}

// ILTriple for IRI Literal sets
func ILTriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	if o == "" {
		return nil
	}

	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewLiteral(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	fmt.Fprintf(b, "%s", qs)
	return err
}

// IBTriple for IRI Literal sets
func IBTriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	if o == "" {
		return nil
	}

	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewBlank(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	fmt.Fprintf(b, "%s", qs)
	return err
}

// BLTriple for IRI Literal sets
func BLTriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	if o == "" {
		return nil
	}

	sub, err := rdf.NewBlank(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewLiteral(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	fmt.Fprintf(b, "%s", qs)
	return err
}

// BITriple for IRI Literal sets
func BITriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	if o == "" {
		return nil
	}

	sub, err := rdf.NewBlank(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewIRI(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	fmt.Fprintf(b, "%s", qs)
	return err
}
