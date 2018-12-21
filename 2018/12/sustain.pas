program Sustain;
{$MODE OBJFPC}

uses
    Sysutils,fgl;

const
    C_INPUT = 'input';
    C_PREFIX = 'initial state: ';

type
    TPots = class
    private
        pots: array of Boolean;
        rules: array[0..31] of Boolean;
        first: Integer;
        last: Integer;
    public
        procedure Evolve;
        function GetString: String;
        procedure Load(var tfInput: TextFile);
        function Sum: Integer;
    private
        procedure LoadRules(var tfInput: TextFile);
        function NextPot(p: Integer): Boolean;
    end;

procedure TPots.Evolve;
var
    e: array of Boolean;
    f, i, l, p: Integer;
begin
    setLength(e, last - first + 5);

    f := 0;
    i := 0;
    l := 0;

    for p := first - 2 to last + 2 do
    begin
        if NextPot(p) then
        begin
            if p < f then f := p;
            e[i] := true;
            l := p;
        end;

        i := i + 1;
    end;

    setLength(pots, l - f + 1);

    i := f - (first - 2);

    for p := 0 to Length(pots) - 1 do pots[p] := e[p + i];

    first := f;
    last := l;
end;

function TPots.GetString: String;
var
    b: Boolean;
    s: String;
begin
    s := '';

    for b in pots do
    begin
        if b then
            s := Concat(s, '#')
        else
            s := Concat(s, '.');
    end;

    GetString := s;
end;

procedure TPots.Load(var tfInput: TextFile);
var
    s: String;
    p: Integer;
begin
    readln(tfInput, s);

    s := Copy(s, Length(C_PREFIX) + 1, Length(s));

    first := 0;
    last := Length(s) - 1;

    setLength(pots, last + 1);

    for p := 0 to last do
    begin
        if s[p + 1] = '#' then
            pots[p] := true
        else
            pots[p] := false;
    end;

    readln(tfInput);

    LoadRules(tfInput);
end;

function TPots.Sum: Integer;
var
    p, s: Integer;
begin
    s := 0;

    for p := first to last do
    begin
        if pots[p - first] then s := s + p;
    end;

    Sum := s;
end;

function TPots.NextPot(p: Integer): Boolean;
var
    i, r: Integer;
begin
    r := 0;

    for i := p - 2 to p + 2 do
    begin
        r := r << 1;

        if (i >= first) and (i <= last) and pots[i - first] then r := r + 1;
    end;

    NextPot := rules[r];
end;

procedure TPots.LoadRules(var tfInput: TextFile);
var
    i, r: Integer;
    s: String;
begin
    while not eof(tfInput) do
    begin
        readln(tfInput, s);

        r := 0;

        for i := 1 to 5 do
        begin
            r := r << 1;
            if s[i] = '#' then r := r + 1;
        end;

        if s[10] = '#' then
            rules[r] := true
        else
            rules[r] := false;
    end;
end;

procedure process(var tfInput: TextFile);
var
    pots: TPots;
    i: Integer;
begin
    pots := TPots.Create;

    try
        pots.Load(tfInput);

        writeln('0: ', pots.GetString);

        for i := 1 to 20 do
        begin
            pots.Evolve;
            writeln(i, ': ', pots.GetString);
        end;

        writeln('Sum: ', pots.Sum);
    finally
        pots.Free;
    end;
end;

procedure main(inputFile: String);
var
    tfInput: TextFile;
begin
    AssignFile(tfInput, inputFile);

    try
        reset(tfInput);

        process(tfInput);

        CloseFile(tfInput);
    except
        on E: EInOutError do
            writeln('Unable to read input. Details: ', E.Message);
    end;
end;

begin
    main(C_INPUT);
end.
